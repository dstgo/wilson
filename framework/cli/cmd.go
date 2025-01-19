package cli

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

func cmdStart(opts *Options) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		configLoader := configureLoader(opts)

		if opts.ConfigFile != "" {
			configLoader = fileLoader(opts)
		}

		startOpts := &StartOptions{
			AppName:      opts.AppName,
			AppVersion:   opts.AppVersion,
			AppBuildTime: opts.AppBuildTime,
			ServiceName:  opts.ServiceName,
			Loader:       configLoader,
		}

		err := opts.StartFn(startOpts)
		if err != nil {
			return fmt.Errorf("failed to start service %s", opts.AppName)
		}
		return nil
	}
}

func cmdVersion(opts *Options) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		cmd.Printf("%s version %s %s/%s\n", opts.AppName, opts.AppVersion, runtime.GOOS, runtime.GOARCH)
		return nil
	}
}
