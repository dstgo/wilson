package circuitbreaker

import (
	"bytes"
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/go-kratos/aegis/circuitbreaker"
	"github.com/go-kratos/aegis/circuitbreaker/sre"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/prometheus/client_golang/prometheus"
	"golang.org/x/exp/rand"

	"github.com/dstgo/wilson/service/gateway/client"
	"github.com/dstgo/wilson/service/gateway/config"
	gtmiddleware "github.com/dstgo/wilson/service/gateway/middleware"
	"github.com/dstgo/wilson/service/gateway/proxy/condition"
	"github.com/dstgo/wilson/service/gateway/utils"
)

func Init(clientFactory client.Factory) {
	breakerFactory := New(clientFactory)
	gtmiddleware.RegisterV2("circuitbreaker", breakerFactory)
	prometheus.MustRegister(_metricDeniedTotal)
}

var (
	_metricDeniedTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "go",
		Subsystem: "gateway",
		Name:      "requests_circuit_breaker_denied_total",
		Help:      "The total number of denied requests",
	}, []string{"protocol", "method", "path", "service", "basePath"})
)

type ratioTrigger struct {
	Ratio int64
	lock  sync.Mutex
	rand  *rand.Rand
}

func newRatioTrigger(ratio int64) *ratioTrigger {
	return &ratioTrigger{
		Ratio: ratio,
		rand:  rand.New(rand.NewSource(uint64(time.Now().UnixNano()))),
	}
}

func (r *ratioTrigger) Allow() error {
	r.lock.Lock()
	defer r.lock.Unlock()
	if r.rand.Int63n(10000) < r.Ratio {
		return nil
	}
	return circuitbreaker.ErrNotAllowed
}
func (*ratioTrigger) MarkSuccess() {}
func (*ratioTrigger) MarkFailed()  {}

type nopTrigger struct{}

func (nopTrigger) Allow() error { return nil }
func (nopTrigger) MarkSuccess() {}
func (nopTrigger) MarkFailed()  {}

func makeBreakerTrigger(in *config.CircuitBreaker) circuitbreaker.CircuitBreaker {
	trigger := in.Trigger
	if trigger != nil && trigger.SuccessRatio != nil {
		var opts []sre.Option
		if trigger.SuccessRatio.Bucket != 0 {
			opts = append(opts, sre.WithBucket(int(trigger.SuccessRatio.Bucket)))
		}
		if trigger.SuccessRatio.Request != 0 {
			opts = append(opts, sre.WithRequest(int64(trigger.SuccessRatio.Request)))
		}
		if trigger.SuccessRatio.Success != 0 {
			opts = append(opts, sre.WithSuccess(trigger.SuccessRatio.Success))
		}
		if trigger.SuccessRatio.Window != 0 {
			opts = append(opts, sre.WithWindow(trigger.SuccessRatio.Window))
		}
		return sre.NewBreaker(opts...)
	}

	if trigger != nil && trigger.Ratio != 0 {
		return newRatioTrigger(trigger.Ratio)
	}

	return nopTrigger{}
}

func makeOnBreakHandler(in *config.CircuitBreaker, factory client.Factory) (http.RoundTripper, io.Closer, error) {
	action := in.Action
	if action.BackupService != nil {
		log.Infof("Making backup service as on break handler: %+v", action.BackupService)
		client, err := factory(&action.BackupService.Endpoint)
		if err != nil {
			return nil, nil, err
		}
		return client, client, nil
	}

	if action.ResponseData != nil {
		log.Infof("Making static response data as on break handler: %+v", action.ResponseData)
		return gtmiddleware.RoundTripperFunc(func(req *http.Request) (*http.Response, error) {
			resp := &http.Response{
				StatusCode: action.ResponseData.StatusCode,
				Header:     http.Header{},
			}
			for _, h := range action.ResponseData.Header {
				resp.Header[h.Key] = h.Value
			}
			resp.Body = io.NopCloser(bytes.NewReader(action.ResponseData.Body))
			return resp, nil
		}), io.NopCloser(nil), nil
	}

	log.Warnf("Unrecoginzed circuit breaker aciton")
	return gtmiddleware.RoundTripperFunc(func(*http.Request) (*http.Response, error) {
		// TBD: on break response
		return &http.Response{
			StatusCode: http.StatusServiceUnavailable,
			Header:     http.Header{},
			Body:       io.NopCloser(&bytes.Buffer{}),
		}, nil
	}), io.NopCloser(nil), nil
}

func isSuccessResponse(conditions []condition.Condition, resp *http.Response) bool {
	return condition.JudgeConditons(conditions, resp, true)
}

func deniedRequestIncr(req *http.Request) {
	labels, ok := gtmiddleware.MetricsLabelsFromContext(req.Context())
	if ok {
		_metricDeniedTotal.WithLabelValues(labels.Protocol(), labels.Method(), labels.Path(), labels.Service(), labels.BasePath()).Inc()
		return
	}
}

func New(factory client.Factory) gtmiddleware.FactoryV2 {
	return func(c *config.Middleware) (gtmiddleware.MiddlewareV2, error) {
		options := &config.CircuitBreaker{}
		if c.Options != nil {
			if err := utils.Copy(c.Options, options); err != nil {
				return nil, err
			}
		}
		breaker := makeBreakerTrigger(options)
		onBreakHandler, closer, err := makeOnBreakHandler(options, factory)
		if err != nil {
			return nil, err
		}
		condtions, err := condition.ParseConditon(options.Conditions)
		if err != nil {
			return nil, err
		}

		return gtmiddleware.NewWithCloser(func(next http.RoundTripper) http.RoundTripper {
			return gtmiddleware.RoundTripperFunc(func(req *http.Request) (*http.Response, error) {
				if err := breaker.Allow(); err != nil {
					// rejected
					// NOTE: when client reject requests locally,
					// continue add counter let the drop ratio higher.
					breaker.MarkFailed()
					deniedRequestIncr(req)
					return onBreakHandler.RoundTrip(req)
				}
				resp, err := next.RoundTrip(req)
				if err != nil {
					breaker.MarkFailed()
					return nil, err
				}
				if !isSuccessResponse(condtions, resp) {
					breaker.MarkFailed()
					return resp, nil
				}
				breaker.MarkSuccess()
				return resp, nil
			})
		}, closer), nil
	}
}
