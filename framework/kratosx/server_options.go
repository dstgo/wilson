package kratosx

import (
	"crypto/tls"

	kmid "github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"

	"github.com/dstgo/wilson/framework/kratosx/config"
	"github.com/dstgo/wilson/framework/kratosx/middleware"
)

func serverOptions(conf config.Config, midOpts []kmid.Middleware) ([]grpc.ServerOption, []http.ServerOption) {
	var gs []grpc.ServerOption
	var hs []http.ServerOption

	// 中间件
	mds := middleware.New(conf)
	mds = append(mds, midOpts...)
	gs = append(gs, grpc.Middleware(mds...))
	hs = append(hs, http.Middleware(mds...))

	// tls
	if conf.App().Server.Tls != nil {
		cert, err := tls.X509KeyPair([]byte(conf.App().Server.Tls.Pem), []byte(conf.App().Server.Tls.Key))
		if err != nil {
			panic(err)
		}
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
			Certificates: []tls.Certificate{cert},
		}

		gs = append(gs, grpc.TLSConfig(tlsConf))
		hs = append(hs, http.TLSConfig(tlsConf))
	}

	return gs, hs
}
