package cli

import (
	kratosConfig "github.com/go-kratos/kratos/v2/config"
	filesource "github.com/go-kratos/kratos/v2/config/file"

	configuresource "github.com/dstgo/wilson/client/rpc/configure"
)

type Loader func() kratosConfig.Source

func fileLoader(opts *Options) Loader {
	return func() kratosConfig.Source {
		return filesource.NewSource(opts.ConfigFile)
	}
}

func configureLoader(opts *Options) Loader {
	return func() kratosConfig.Source {
		return configuresource.New(opts.ConfigHost, opts.ConfigToken, opts.AppName)
	}
}
