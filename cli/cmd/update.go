package cmd

import (
	"github.com/nednella/bootstrap.sh/internal/jobs"
	"github.com/nednella/bootstrap.sh/internal/ui"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:     "update",
	Short:   "Update binary and pull latest content",
	GroupID: jobsGroupID,
	Run: func(cmd *cobra.Command, args []string) {
		ui.Header("update")
		err := jobs.Update()
		if err != nil {
			ui.Die(err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
