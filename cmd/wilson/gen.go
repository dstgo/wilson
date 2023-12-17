package main

import (
	"github.com/dstgo/filebox"
	"github.com/dstgo/wilson/assets"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"log/slog"
)

var genDir string
var genCmd = &cobra.Command{
	Use:     "gen [--d dir]",
	Short:   "Generate the default wilson config directory",
	Long:    "Generate the default wilson config directory. if already exists, all files will be overwritten.",
	Example: "wilson gen --d /etc/wilson",
	Run: func(cmd *cobra.Command, args []string) {
		err := generateResourceDir(genDir)
		if err != nil {
			slog.Error(err.Error())
		}
	},
}

func init() {
	genCmd.Flags().StringVar(&genDir, "d", DefaultDir, "default wilson resource directory")
}

func generateResourceDir(dir string) error {
	err := filebox.CopyFsDir(assets.Fs, filebox.Os, "config", dir, filebox.DefaultBuffer)
	if err != nil {
		return errors.Wrap(err, "generate config failed")
	}
	return nil
}
