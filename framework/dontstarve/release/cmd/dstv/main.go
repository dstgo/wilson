package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:           "dstv",
	Short:         "dstv is a cmd tool for fetching DST (Don't Starve Together) game updates",
	SilenceUsage:  true,
	SilenceErrors: true,
}

func init() {
	rootCmd.AddCommand(latestCmd)
	rootCmd.AddCommand(listCmd)
}

func main() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
