package cmd

import (
	"github.com/nednella/bootstrap.sh/internal/jobs"
	"github.com/nednella/bootstrap.sh/internal/ui"
	"github.com/spf13/cobra"
)

var macosCmd = &cobra.Command{
	Use:     "macos",
	Short:   "Apply macOS preferences",
	GroupID: jobsGroupID,
	Run: func(cmd *cobra.Command, args []string) {
		ui.Header("macos")
		err := jobs.MacOS()
		if err != nil {
			ui.Die(err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(macosCmd)
}
