package cli

import (
	"errors"
	"fmt"
	"runtime"

	kratosConfig "github.com/go-kratos/kratos/v2/config"
	"github.com/spf13/cobra"
)

func newVersionCmd(opts *Options) *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print the service version",
		RunE:  runVersionCmd(opts),
	}
}

func newStartCmd(opts *Options) *cobra.Command {
	return &cobra.Command{
		Use:   "start",
		Short: "Start the service",
		RunE:  runStartCmd(opts),
	}
}

func runVersionCmd(opts *Options) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		cmd.Printf("%s version %s %s/%s\n", opts.AppName, opts.AppVersion, runtime.GOOS, runtime.GOARCH)
		return nil
	}
}

func runStartCmd(opts *Options) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {

		var configLoader func() kratosConfig.Source

		if opts.ConfigHost != "" && opts.ConfigToken != "" {
			configLoader = configureLoader(opts)
		}

		if opts.ConfigFile != "" {
			configLoader = fileLoader(opts)
		}

		if configLoader == nil {
			return errors.New("no config source provided")
		}

		startOpts := &StartOptions{
			AppName:    opts.AppName,
			AppVersion: opts.AppVersion,
			ServiceID:  opts.ServiceID,
			Loader:     configLoader,
		}

		err := opts.StartFn(startOpts)
		if err != nil {
			return fmt.Errorf("failed to start service %s: %s", opts.AppName, err)
		}
		return nil
	}
}
