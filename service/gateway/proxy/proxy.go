package proxy

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync/atomic"
	"time"

	"github.com/go-kratos/aegis/circuitbreaker/sre"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/selector"
	"github.com/go-kratos/kratos/v2/transport/http/status"
	"github.com/prometheus/client_golang/prometheus"

	"github.com/dstgo/wilson/framework/constants"
	"github.com/dstgo/wilson/service/gateway/client"
	"github.com/dstgo/wilson/service/gateway/config"
	gtmiddleware "github.com/dstgo/wilson/service/gateway/middleware"
	"github.com/dstgo/wilson/service/gateway/router"
	"github.com/dstgo/wilson/service/gateway/router/mux"
)

var (
	_metricRequestsTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "go",
		Subsystem: "gateway",
		Name:      "requests_code_total",
		Help:      "The total number of processed requests",
	}, []string{"protocol", "method", "path", "code", "service", "basePath"})
	_metricRequestsDuration = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "go",
		Subsystem: "gateway",
		Name:      "requests_duration_seconds",
		Help:      "Requests duration(sec).",
		Buckets:   []float64{0.005, 0.01, 0.025, 0.05, 0.1, 0.250, 0.5, 1},
	}, []string{"protocol", "method", "path", "service", "basePath"})
	_metricSentBytes = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "go",
		Subsystem: "gateway",
		Name:      "requests_tx_bytes",
		Help:      "Total sent connection bytes",
	}, []string{"protocol", "method", "path", "service", "basePath"})
	_metricReceivedBytes = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "go",
		Subsystem: "gateway",
		Name:      "requests_rx_bytes",
		Help:      "Total received connection bytes",
	}, []string{"protocol", "method", "path", "service", "basePath"})
	_metricRetryState = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "go",
		Subsystem: "gateway",
		Name:      "requests_retry_state",
		Help:      "Total request retries",
	}, []string{"protocol", "method", "path", "service", "basePath", "success"})
)

func init() {
	prometheus.MustRegister(_metricRequestsTotal)
	prometheus.MustRegister(_metricRequestsDuration)
	prometheus.MustRegister(_metricRetryState)
	prometheus.MustRegister(_metricSentBytes)
	prometheus.MustRegister(_metricReceivedBytes)
}

func setXFFHeader(req *http.Request) {
	// see https://github.com/golang/go/blob/master/src/net/http/httputil/reverseproxy.go
	if clientIP, _, err := net.SplitHostPort(req.RemoteAddr); err == nil {
		// If we aren't the first proxy retain prior
		// X-Forwarded-For information as a comma+space
		// separated list and fold multiple headers into one.
		prior, ok := req.Header["X-Forwarded-For"]
		omit := ok && prior == nil // Issue 38079: nil now means don't populate the header
		if len(prior) > 0 {
			clientIP = strings.Join(prior, ", ") + ", " + clientIP
		}
		if !omit {
			req.Header.Set("X-Forwarded-For", clientIP)
		}
	}
}

func writeError(w http.ResponseWriter, r *http.Request, err error, labels gtmiddleware.MetricsLabels) {
	var statusCode int
	switch {
	case errors.Is(err, context.Canceled),
		err.Error() == "client disconnected":
		statusCode = 499
	case errors.Is(err, context.DeadlineExceeded):
		statusCode = 504
	default:
		log.Errorf("Failed to handle request: %s: %+v", r.URL.String(), err)
		statusCode = 502
	}
	requestsTotalIncr(labels, statusCode)
	if labels.Protocol() == constants.GRPC {
		// see https://github.com/googleapis/googleapis/blob/master/google/rpc/code.proto
		code := strconv.Itoa(int(status.ToGRPCCode(statusCode)))
		w.Header().Set("Content-Type", "application/grpc")
		w.Header().Set("Grpc-Status", code)
		w.Header().Set("Grpc-Message", err.Error())
		statusCode = 200
	}
	w.WriteHeader(statusCode)
}

// notFoundHandler replies to the request with an HTTP 404 not found error.
func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	code := http.StatusNotFound
	message := "404 page not found"
	http.Error(w, message, code)
	log.Context(r.Context()).Errorw(
		"source", "accesslog",
		"host", r.Host,
		"method", r.Method,
		"path", r.URL.Path,
		"query", r.URL.RawQuery,
		"user_agent", r.Header.Get("User-Agent"),
		"code", code,
		"error", message,
	)
	_metricRequestsTotal.WithLabelValues(constants.HTTP, r.Method, "/404", strconv.Itoa(code), "", "").Inc()
}

func methodNotAllowedHandler(w http.ResponseWriter, r *http.Request) {
	code := http.StatusMethodNotAllowed
	message := http.StatusText(code)
	http.Error(w, message, code)
	log.Context(r.Context()).Errorw(
		"source", "accesslog",
		"host", r.Host,
		"method", r.Method,
		"path", r.URL.Path,
		"query", r.URL.RawQuery,
		"user_agent", r.Header.Get("User-Agent"),
		"code", code,
		"error", message,
	)
	_metricRequestsTotal.WithLabelValues(constants.HTTP, r.Method, "/405", strconv.Itoa(code), "", "").Inc()
}

type interceptors struct {
	prepareAttemptTimeoutContext func(ctx context.Context, req *http.Request, timeout time.Duration) (context.Context, context.CancelFunc)
}

func (i *interceptors) SetPrepareAttemptTimeoutContext(f func(ctx context.Context, req *http.Request, timeout time.Duration) (context.Context, context.CancelFunc)) {
	if f != nil {
		i.prepareAttemptTimeoutContext = f
	}
}

// Proxy is a gateway proxy.
type Proxy struct {
	router            atomic.Value
	clientFactory     client.Factory
	Interceptors      interceptors
	middlewareFactory gtmiddleware.FactoryV2
}

// New is new a gateway proxy.
func New(clientFactory client.Factory, middlewareFactory gtmiddleware.FactoryV2) (*Proxy, error) {
	p := &Proxy{
		clientFactory:     clientFactory,
		middlewareFactory: middlewareFactory,
		Interceptors: interceptors{
			prepareAttemptTimeoutContext: defaultAttemptTimeoutContext,
		},
	}
	p.router.Store(mux.NewRouter(http.HandlerFunc(notFoundHandler), http.HandlerFunc(methodNotAllowedHandler)))
	return p, nil
}

func (p *Proxy) buildMiddleware(ms []config.Middleware, next http.RoundTripper) (http.RoundTripper, error) {
	for i := len(ms) - 1; i >= 0; i-- {
		m, err := p.middlewareFactory(&ms[i])
		if err != nil {
			if errors.Is(err, gtmiddleware.ErrNotFound) {
				log.Errorf("Skip does not exist middleware: %s", ms[i].Name)
				continue
			}
			return nil, err
		}
		next = m.Process(next)
	}
	return next, nil
}

func splitRetryMetricsHandler(e *config.Endpoint) (func(int), func(int, error)) {
	labels := gtmiddleware.NewMetricsLabels(e)
	success := func(i int) {
		if i <= 0 {
			return
		}
		retryStateIncr(labels, true)
	}
	failed := func(i int, err error) {
		if i <= 0 {
			return
		}
		if errors.Is(err, context.Canceled) {
			return
		}
		retryStateIncr(labels, false)
	}
	return success, failed
}

func (p *Proxy) buildEndpoint(e *config.Endpoint, ms []config.Middleware) (_ http.Handler, _ io.Closer, retError error) {
	client, err := p.clientFactory(e)
	if err != nil {
		return nil, nil, err
	}
	tripper := http.RoundTripper(client)
	closer := io.Closer(client)
	defer closeOnError(closer, &retError)

	tripper, err = p.buildMiddleware(e.Middlewares, tripper)
	if err != nil {
		return nil, nil, err
	}
	tripper, err = p.buildMiddleware(ms, tripper)
	if err != nil {
		return nil, nil, err
	}
	retryStrategy, err := prepareRetryStrategy(e)
	if err != nil {
		return nil, nil, err
	}
	labels := gtmiddleware.NewMetricsLabels(e)
	markSuccessStat, markFailedStat := splitRetryMetricsHandler(e)
	retryBreaker := sre.NewBreaker(sre.WithSuccess(0.8))
	markSuccess := func(i int) {
		markSuccessStat(i)
		if i > 0 {
			retryBreaker.MarkSuccess()
		}
	}
	markFailed := func(i int, err error) {
		markFailedStat(i, err)
		if i > 0 {
			retryBreaker.MarkFailed()
		}
	}
	return http.Handler(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		startTime := time.Now()
		setXFFHeader(req)

		reqOpts := gtmiddleware.NewRequestOptions(e)
		ctx := gtmiddleware.NewRequestContext(req.Context(), reqOpts)
		ctx, cancel := context.WithTimeout(ctx, retryStrategy.timeout)
		defer cancel()
		defer func() {
			requestsDurationObserve(labels, time.Since(startTime).Seconds())
		}()

		body, err := io.ReadAll(req.Body)
		if err != nil {
			writeError(w, req, err, labels)
			return
		}
		receivedBytesAdd(labels, int64(len(body)))
		req.GetBody = func() (io.ReadCloser, error) {
			reader := bytes.NewReader(body)
			return io.NopCloser(reader), nil
		}

		var resp *http.Response
		for i := 0; i < retryStrategy.attempts; i++ {
			if i > 0 {
				if !retryFeature.Enabled() {
					break
				}
				if err := retryBreaker.Allow(); err != nil {
					markFailed(i, err)
					break
				}
			}

			if (i + 1) >= retryStrategy.attempts {
				reqOpts.LastAttempt = true
			}
			// canceled or deadline exceeded
			if err = ctx.Err(); err != nil {
				markFailed(i, err)
				break
			}
			tryCtx, cancel := p.Interceptors.prepareAttemptTimeoutContext(ctx, req, retryStrategy.perTryTimeout)
			defer cancel()
			reader := bytes.NewReader(body)
			req.Body = io.NopCloser(reader)
			resp, err = tripper.RoundTrip(req.Clone(tryCtx))
			if err != nil {
				markFailed(i, err)
				log.Errorf("Attempt at [%d/%d], failed to handle request: %s: %+v", i+1, retryStrategy.attempts, req.URL.String(), err)
				continue
			}
			if !judgeRetryRequired(retryStrategy.conditions, resp) {
				reqOpts.LastAttempt = true
				markSuccess(i)
				break
			}
			markFailed(i, errors.New("assertion failed"))
			// continue the retry loop
		}
		if err != nil {
			writeError(w, req, err, labels)
			return
		}

		hasJson := strings.Contains(resp.Header.Get("Content-Type"), "json")
		hasEvent := strings.Contains(resp.Header.Get("Content-Type"), "text/event-stream")
		headers := w.Header()
		for k, v := range resp.Header {
			if k == "Content-Length" && hasJson {
				continue
			}
			headers[k] = v
		}
		w.WriteHeader(resp.StatusCode)

		doCopyBody := func() bool {
			if resp.Body == nil {
				return true
			}
			defer resp.Body.Close()

			eventCopy := func(dst io.Writer, src io.Reader) (int64, error) {
				flusher, ok := w.(http.Flusher)
				if !ok {
					return io.Copy(dst, src)
				}

				buf := make([]byte, 1024)
				written := int64(0)
				for {
					nr, er := src.Read(buf)
					if nr > 0 {
						nw, ew := dst.Write(buf[0:nr])
						if nw < 0 || nr < nw {
							if ew == nil {
								return 0, ew
							}
						}
						written += int64(nw)
						if ew != nil {
							return 0, ew
						}
						if nr != nw {
							return 0, io.ErrShortWrite
						}
					}
					flusher.Flush()

					if er != nil && er != io.EOF {
						return 0, er
					}
					if er == io.EOF {
						break
					}
				}
				return written, nil
			}

			var sent int64
			if reqOpts.Endpoint.ResponseFormat && hasJson {
				body := ResponseFormat(resp)
				sent = int64(len(body))
				_, err = w.Write(body)
			} else if hasEvent {
				sent, err = eventCopy(w, resp.Body)
			} else {
				sent, err = io.Copy(w, resp.Body)
			}

			resp.Header.Set("Content-Length", fmt.Sprint(sent))

			if err != nil {
				reqOpts.DoneFunc(ctx, selector.DoneInfo{Err: err})
				sentBytesAdd(labels, sent)
				log.Errorf("Failed to copy backend response body to client: [%s] %s %s %d %+v\n", e.Protocol, e.Method, e.Path, sent, err)
				return false
			}
			sentBytesAdd(labels, sent)
			reqOpts.DoneFunc(ctx, selector.DoneInfo{ReplyMD: resp.Trailer})
			// see https://pkg.go.dev/net/http#example-ResponseWriter-Trailers
			for k, v := range resp.Trailer {
				headers[http.TrailerPrefix+k] = v
			}
			return true
		}
		doCopyBody()
		requestsTotalIncr(labels, resp.StatusCode)
	})), closer, nil
}

func receivedBytesAdd(labels gtmiddleware.MetricsLabels, received int64) {
	_metricReceivedBytes.WithLabelValues(labels.Protocol(), labels.Method(), labels.Path(), labels.Service(), labels.BasePath()).Add(float64(received))
}

func sentBytesAdd(labels gtmiddleware.MetricsLabels, sent int64) {
	_metricSentBytes.WithLabelValues(labels.Protocol(), labels.Method(), labels.Path(), labels.Service(), labels.BasePath()).Add(float64(sent))
}

func requestsTotalIncr(labels gtmiddleware.MetricsLabels, statusCode int) {
	_metricRequestsTotal.WithLabelValues(labels.Protocol(), labels.Method(), labels.Path(), strconv.Itoa(statusCode), labels.Service(), labels.BasePath()).Inc()
}

func requestsDurationObserve(labels gtmiddleware.MetricsLabels, seconds float64) {
	_metricRequestsDuration.WithLabelValues(labels.Protocol(), labels.Method(), labels.Path(), labels.Service(), labels.BasePath()).Observe(seconds)
}

func retryStateIncr(labels gtmiddleware.MetricsLabels, success bool) {
	if success {
		_metricRetryState.WithLabelValues(labels.Protocol(), labels.Method(), labels.Path(), labels.Service(), labels.BasePath(), "true").Inc()
		return
	}
	_metricRetryState.WithLabelValues(labels.Protocol(), labels.Method(), labels.Path(), labels.Service(), labels.BasePath(), "false").Inc()
}

func closeOnError(closer io.Closer, err *error) {
	if *err == nil {
		return
	}
	_ = closer.Close()
}

// Update updates service endpoint.
func (p *Proxy) Update(c *config.Config) (retError error) {
	router := mux.NewRouter(http.HandlerFunc(notFoundHandler), http.HandlerFunc(methodNotAllowedHandler))
	for _, e := range c.Endpoints {
		ep := e
		handler, closer, err := p.buildEndpoint(&ep, c.Middlewares)
		if err != nil {
			return err
		}
		defer closeOnError(closer, &retError)
		if err = router.Handle(e.Path, e.Method, e.Host, handler, closer); err != nil {
			return err
		}
		log.Infof("build endpoint: [%s] %s %s", e.Protocol, e.Method, e.Path)
	}
	old := p.router.Swap(router)
	tryCloseRouter(old)
	return nil
}

func tryCloseRouter(in any) {
	if in == nil {
		return
	}
	r, ok := in.(router.Router)
	if !ok {
		return
	}
	go func() {
		ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
		defer cancel()
		_ = r.SyncClose(ctx)
	}()
}

func (p *Proxy) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			w.WriteHeader(http.StatusBadGateway)
			buf := make([]byte, 64<<10) //nolint:gomnd
			n := runtime.Stack(buf, false)
			log.Errorf("panic recovered: %s", buf[:n])
			fmt.Fprintf(os.Stderr, "panic recovered: %s\n", buf[:n])
		}
	}()

	srv, ok := p.router.Load().(router.Router)
	if !ok {
		panic(fmt.Sprintf("invalid type, expected router.Router, got %T", p.router))
	}

	srv.ServeHTTP(w, req)
}

// DebugHandler implemented debug handler.
func (p *Proxy) DebugHandler() http.Handler {
	debugMux := http.NewServeMux()
	debugMux.HandleFunc("/debug/proxy/router/inspect", func(rw http.ResponseWriter, r *http.Request) {
		router, ok := p.router.Load().(router.Router)
		if !ok {
			return
		}
		inspect := mux.InspectMuxRouter(router)
		rw.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(rw).Encode(inspect)
	})
	return debugMux
}
