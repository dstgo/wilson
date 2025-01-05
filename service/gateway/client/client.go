package client

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/selector"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"

	"github.com/dstgo/wilson/service/gateway/consts"
	"github.com/dstgo/wilson/service/gateway/middleware"
)

type client struct {
	applier  *nodeApplier
	selector selector.Selector
}

type Client interface {
	http.RoundTripper
	io.Closer
}

func newClient(applier *nodeApplier, selector selector.Selector) *client {
	return &client{
		applier:  applier,
		selector: selector,
	}
}

func (c *client) Close() error {
	c.applier.Cancel()
	return nil
}

func (c *client) RoundTrip(req *http.Request) (*http.Response, error) {
	ctx := req.Context()
	reqOpt, _ := middleware.FromRequestContext(ctx)
	filter, _ := middleware.SelectorFiltersFromContext(ctx)
	n, done, err := c.selector.Select(ctx, selector.WithNodeFilter(filter...))
	if err != nil {
		return nil, err
	}
	reqOpt.CurrentNode = n

	addr := n.Address()
	reqOpt.Backends = append(reqOpt.Backends, addr)
	req.URL.Host = addr
	req.URL.Scheme = "http"
	req.RequestURI = ""
	startAt := time.Now()

	// Inject the context into the HTTP headers
	otel.GetTextMapPropagator().Inject(ctx, propagation.HeaderCarrier(req.Header))

	resp, err := n.(*node).client.Do(req)
	reqOpt.UpstreamResponseTime = append(reqOpt.UpstreamResponseTime, time.Since(startAt).Seconds())
	if err != nil {
		done(ctx, selector.DoneInfo{Err: err})
		reqOpt.UpstreamStatusCode = append(reqOpt.UpstreamStatusCode, 0)
		return nil, err
	}

	val := tracing.TraceID()(ctx)
	traceId, ok := val.(string)
	if !ok {
		panic(fmt.Sprintf("invalid traceId type, expected string, got %T", val))
	}
	resp.Header.Set(consts.TRACE_ID, traceId)

	reqOpt.UpstreamStatusCode = append(reqOpt.UpstreamStatusCode, resp.StatusCode)
	reqOpt.DoneFunc = done
	return resp, nil
}
