package configure

import (
	"context"
	"crypto/tls"
	"crypto/x509"

	"github.com/go-kratos/kratos/v2/selector"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	ggrpc "google.golang.org/grpc"
)

type client struct {
	applier  *nodeApplier
	selector selector.Selector
}

type Client interface {
	Conn(ctx context.Context) (*ggrpc.ClientConn, error)
}

func newClient(applier *nodeApplier, selector selector.Selector) Client {
	return &client{
		applier:  applier,
		selector: selector,
	}
}

func (c *client) Close() error {
	return nil
}

func (c *client) Conn(ctx context.Context) (*ggrpc.ClientConn, error) {
	if c.applier.endpoint.Type == DIRECT {
		return c.connByDirect(ctx)
	}
	return c.connByDiscovery(ctx)
}

func (c *client) connByDirect(ctx context.Context) (*ggrpc.ClientConn, error) {
	n, done, err := c.selector.Select(ctx)
	if err != nil {
		return nil, err
	}
	opts := c.options()
	opts = append(opts,
		grpc.WithEndpoint(n.Address()),
	)
	conn, err := grpc.DialInsecure(ctx, opts...)
	if err != nil {
		done(ctx, selector.DoneInfo{Err: err})
		return nil, err
	}
	return conn, nil
}

func (c *client) connByDiscovery(ctx context.Context) (*ggrpc.ClientConn, error) {
	opts := c.options()
	opts = append(opts,
		grpc.WithEndpoint(DISCOVERY+":///"+c.applier.endpoint.Server),
		grpc.WithDiscovery(c.applier.registry),
	)
	return grpc.DialInsecure(ctx, opts...)
}

func (c *client) options() []grpc.ClientOption {
	var opts = []grpc.ClientOption{
		grpc.WithMiddleware(Middlewares(c.applier.endpoint)...),
		grpc.WithTimeout(c.applier.endpoint.Timeout),
	}

	// tls
	if c.applier.endpoint.Tls != nil {
		cp := x509.NewCertPool()
		if cp.AppendCertsFromPEM([]byte(c.applier.endpoint.Tls.Ca)) {
			tlsConf := &tls.Config{
				MinVersion: tls.VersionTLS12,
				CipherSuites: []uint16{
					tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
					tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
					tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,
					tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,
					tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
					tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
				},
				CurvePreferences: []tls.CurveID{
					tls.X25519,
					tls.CurveP256,
				},
				ServerName: c.applier.endpoint.Tls.Name,
				RootCAs:    cp,
			}
			opts = append(opts, grpc.WithTLSConfig(tlsConf))
		}
	}
	return opts
}
