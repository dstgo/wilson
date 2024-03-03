package main

import (
	"github.com/spf13/cobra"
)

var (
	Author    string
	Version   string
	BuildTime string
)

var rootCmd = &cobra.Command{
	Use:  "wigfrid [command]",
	Long: `wigfrid is the daemon of the wendy panel, use wigfrid command to build a local dst containers manager.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}

func main() {
	rootCmd.Execute()
}
