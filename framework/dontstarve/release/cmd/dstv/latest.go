package main

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/dstgo/wilson/framework/dontstarve/release"
)

var unstable bool

var latestCmd = &cobra.Command{
	Use:   "latest",
	Short: "fetch the latest version",
	RunE: func(cmd *cobra.Command, args []string) error {
		version, err := FetchLatestVersion(!unstable)
		if err != nil {
			return err
		}
		fmt.Printf("%d\n", version)
		return nil
	},
}

func init() {
	latestCmd.Flags().BoolVar(&unstable, "unstable", false, "fetch unstable version")
}

func FetchLatestVersion(unstable bool) (int, error) {
	latest, err := release.New().Latest(!unstable)
	if err != nil {
		return 0, err
	}
	return latest.Version, nil
}
