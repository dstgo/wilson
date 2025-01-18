package tracing

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/go-kratos/kratos/v2"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"go.opentelemetry.io/otel/trace"

	"github.com/dstgo/wilson/service/gateway/config"
	gtmiddleware "github.com/dstgo/wilson/service/gateway/middleware"
	"github.com/dstgo/wilson/service/gateway/utils"
)

const (
	defaultTimeout     = time.Duration(10 * time.Second)
	defaultServiceName = "gateway"
	defaultTracerName  = "gateway"
)

var globaltp = &struct {
	provider trace.TracerProvider
	initOnce sync.Once
}{}

func init() {
	gtmiddleware.Register("tracing", Middleware)
}

// Middleware is a opentelemetry middleware.
func Middleware(c *config.Middleware) (gtmiddleware.Middleware, error) {
	options := &config.Tracing{}
	if c.Options != nil {
		if err := utils.Copy(c.Options, options); err != nil {
			return nil, err
		}
	}
	if globaltp.provider == nil {
		globaltp.initOnce.Do(func() {
			globaltp.provider = newTracerProvider(context.Background(), options)
			propagator := propagation.NewCompositeTextMapPropagator(propagation.Baggage{}, propagation.TraceContext{})
			otel.SetTracerProvider(globaltp.provider)
			otel.SetTextMapPropagator(propagator)
		})
	}
	tracer := otel.Tracer(defaultTracerName)
	return func(next http.RoundTripper) http.RoundTripper {
		return gtmiddleware.RoundTripperFunc(func(req *http.Request) (reply *http.Response, err error) {
			ctx, span := tracer.Start(
				req.Context(),
				fmt.Sprintf("%s %s", req.Method, req.URL.Path),
				trace.WithSpanKind(trace.SpanKindClient),
			)

			// attributes for each request
			span.SetAttributes(
				semconv.HTTPMethodKey.String(req.Method),
				semconv.HTTPTargetKey.String(req.URL.Path),
				semconv.NetPeerIPKey.String(req.RemoteAddr),
			)

			defer func() {
				if err != nil {
					span.RecordError(err)
					span.SetStatus(codes.Error, err.Error())
				} else {
					span.SetStatus(codes.Ok, "OK")
				}
				if reply != nil {
					span.SetAttributes(semconv.HTTPStatusCodeKey.Int(reply.StatusCode))
				}
				span.End()
			}()
			return next.RoundTrip(req.WithContext(ctx))
		})
	}, nil
}

func newTracerProvider(ctx context.Context, options *config.Tracing) trace.TracerProvider {
	var (
		timeout     = defaultTimeout
		serviceName = defaultServiceName
	)

	if appInfo, ok := kratos.FromContext(ctx); ok {
		serviceName = appInfo.Name()
	}

	if options.Timeout != 0 {
		timeout = options.Timeout
	}

	var sampler sdktrace.Sampler
	if options.SampleRatio == 0 {
		sampler = sdktrace.AlwaysSample()
	} else {
		sampler = sdktrace.TraceIDRatioBased(options.SampleRatio)
	}

	// attributes for all requests
	resources := resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceNameKey.String(serviceName),
	)

	providerOptions := []sdktrace.TracerProviderOption{
		sdktrace.WithSampler(sampler),
		sdktrace.WithResource(resources),
	}

	otlpoptions := []otlptracehttp.Option{}
	if options.Endpoint != "" {
		otlpoptions = append(otlpoptions, otlptracehttp.WithEndpoint(options.Endpoint))

		if options.Timeout != 0 {
			otlpoptions = append(otlpoptions, otlptracehttp.WithTimeout(timeout))
		}

		if !options.Insecure {
			otlpoptions = append(otlpoptions, otlptracehttp.WithInsecure())
		}

		client := otlptracehttp.NewClient(
			otlpoptions...,
		)

		exporter, err := otlptrace.New(ctx, client)
		if err != nil {
			log.Fatalf("creating OTLP trace exporter: %v", err)
		}
		providerOptions = append(providerOptions, sdktrace.WithBatcher(exporter))
	}

	return sdktrace.NewTracerProvider(providerOptions...)
}
