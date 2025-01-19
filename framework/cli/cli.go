package cli

import (
	"runtime"
	"time"

	"github.com/pkg/errors"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
)

type Options struct {
	AppName      string
	AppVersion   string
	AppBuildTime string
	ConfigFile   string
	Description  string

	StartFn func(opt *Options) error
}

func NewCLI(opts *Options) Service {

	opts.AppVersion = lo.Ternary(opts.AppVersion == "", "v0.0.0", opts.AppVersion)
	opts.AppBuildTime = lo.Ternary(opts.AppBuildTime == "", time.Time{}.Format(time.DateTime), opts.AppBuildTime)

	rootCmd := &cobra.Command{
		Use:           opts.AppName,
		SilenceUsage:  true,
		SilenceErrors: true,
		Short:         opts.Description,
	}

	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "Print the service version",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Printf("%s version %s %s/%s\n", opts.AppName, opts.AppVersion, runtime.GOOS, runtime.GOARCH)
		},
	}

	startCmd := &cobra.Command{
		Use:   "start",
		Short: "Start the service",
		RunE: func(cmd *cobra.Command, args []string) error {
			err := opts.StartFn(opts)
			if err != nil {
				return errors.Wrapf(err, "service %s failed to start", opts.AppName)
			}
			return nil
		},
	}

	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(startCmd)

	return Service{
		opts: opts,
		cmd:  rootCmd,
	}
}

type Service struct {
	opts *Options
	cmd  *cobra.Command
}

func (c Service) Parse() {
	if c.cmd != nil {
		c.cmd.PersistentFlags().StringVarP(&c.opts.ConfigFile, "conf", "f", "conf.yaml", "config file path, default: conf.yaml")
	}
}

func (c Service) Start() {
	err := c.cmd.Execute()
	if err != nil {
		c.cmd.PrintErrf("Error: %v\n", err)
	}
}

func (c Service) CMD() *cobra.Command {
	return c.cmd
}
