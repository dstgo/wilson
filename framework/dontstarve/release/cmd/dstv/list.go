package main

import (
	"cmp"
	"fmt"

	"github.com/spf13/cobra"
	"golang.org/x/exp/slices"

	"github.com/dstgo/wilson/framework/dontstarve/release"
)

var page int

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list available versions",
	RunE: func(cmd *cobra.Command, args []string) error {
		list, err := ListVersion(page)
		if err != nil {
			return err
		}

		slices.SortFunc(list, func(a, b release.Release) int {
			return -cmp.Compare(a.Version, b.Version)
		})

		for _, r := range list {
			if r.Pinned {
				fmt.Printf("%d\t%s\tPinned\n", r.Version, r.Branch)
			} else {
				fmt.Printf("%d\t%s\n", r.Version, r.Branch)
			}
		}
		return nil
	},
}

func init() {
	listCmd.Flags().IntVar(&page, "page", 1, "specify the page number, default to 1")
}

func ListVersion(page int) ([]release.Release, error) {
	list, err := release.New().List(page)
	if err != nil {
		return nil, err
	}
	return list, nil
}
