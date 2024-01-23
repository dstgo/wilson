package main

import (
	"github.com/spf13/cobra"
)

const (
	DefaultDir = "/etc/wilson"

	DataDir = "/var/lib/wilson"
)

var rootCmd = &cobra.Command{
	Use:   "wilson [command]",
	Short: "Wilson is the backend web server of wendy panel",
	Long: `Wilson is the backend web server of wendy panel. If you are starting it for the first time,
it is recommended to use the gen command to initialize the wilson configuration directory,
then complete the configuration file according to your needs. 

Access https://github.com/dstgo/wilson for more information.`,
}

var (
	Author    string
	Version   string
	BuildTime string
)

func init() {
	// subcommands
	rootCmd.AddCommand(genCmd)
	rootCmd.AddCommand(serverCmd)
}

func main() {
	rootCmd.Execute()
}
