package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	Author    string
	Version   string
	BuildTime string
)

var rootCmd = &cobra.Command{
	Use: "wig [command]",
	Long: `wigfrid is the daemon of the wendy panel, use wig command to build a local dst containers manager, 
go https://github.com/dstgo/wigfrid to see more information.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
