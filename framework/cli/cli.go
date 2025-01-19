package cli

import (
	"fmt"
	"os"
	"time"
	_ "time/tzdata"

	"github.com/samber/lo"
	"github.com/spf13/cobra"
)

type Options struct {
	AppName      string
	AppVersion   string
	AppBuildTime string
	Description  string

	ConfigFile  string
	ConfigHost  string
	ConfigToken string
	ServiceName string
	Timezone    string

	StartFn func(opt *StartOptions) error
}

type StartOptions struct {
	Loader Loader

	AppName      string
	AppVersion   string
	AppBuildTime string
	ServiceName  string
	Timezone     string
}

func NewCLI(opts *Options) Service {
	if opts == nil {
		opts = &Options{}
	}

	if opts.AppName == "" {
		panic("app name is required")
	}

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
		RunE:  cmdVersion(opts),
	}

	startCmd := &cobra.Command{
		Use:   "start",
		Short: "Start the service",
		RunE:  cmdStart(opts),
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

const (
	EnvConfigFile  = "CONF_FILE"
	EnvConfigHost  = "CONF_HOST"
	EnvConfigToken = "CONF_TOKEN"
	EnvServiceName = "SERVICE_NAME"
	EnvTimezone    = "TIMEZONE"
)

func (c Service) Parse() {
	if c.cmd != nil {
		c.cmd.PersistentFlags().StringVar(&c.opts.ConfigFile, "file", "", "config file path")
		c.cmd.PersistentFlags().StringVar(&c.opts.ConfigHost, "host", "", "configure host")
		c.cmd.PersistentFlags().StringVar(&c.opts.ConfigToken, "token", "", "configure token")
		c.cmd.PersistentFlags().StringVar(&c.opts.ServiceName, "name", "", "service name")
		c.cmd.PersistentFlags().StringVar(&c.opts.Timezone, "timezone", "", "time zone")

		if c.opts.ConfigFile == "" {
			c.opts.ConfigFile = os.Getenv(EnvConfigFile)
		}
		if c.opts.ConfigHost == "" {
			c.opts.ConfigHost = os.Getenv(EnvConfigHost)
		}
		if c.opts.ConfigToken == "" {
			c.opts.ConfigToken = os.Getenv(EnvConfigToken)
		}
		if c.opts.ServiceName == "" {
			c.opts.ServiceName = os.Getenv(EnvServiceName)
		}
		if c.opts.Timezone == "" {
			c.opts.Timezone = os.Getenv(EnvTimezone)
		}
	}
}

func (c Service) prepare() error {
	if c.opts.Timezone == "" {
		c.opts.Timezone = "Local"
	}

	location, err := time.LoadLocation(c.opts.Timezone)
	if err != nil {
		return fmt.Errorf("invalid timezone: %s", c.opts.Timezone)
	}
	time.Local = location

	return nil
}

func (c Service) Start() {
	err := c.prepare()
	if err != nil {
		c.cmd.PrintErrf("Error: %v\n", err)
		return
	}

	err = c.cmd.Execute()
	if err != nil {
		c.cmd.PrintErrf("Error: %v\n", err)
	}
}

func (c Service) CMD() *cobra.Command {
	return c.cmd
}
