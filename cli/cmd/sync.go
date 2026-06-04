package cmd

import (
	"github.com/nednella/bootstrap.sh/internal/jobs"
	"github.com/nednella/bootstrap.sh/internal/ui"
	"github.com/spf13/cobra"
)

var syncCmd = &cobra.Command{
	Use:     "sync",
	Short:   "Pull the latest changes from the remote repository",
	GroupID: jobsGroupID,
	Run: func(cmd *cobra.Command, args []string) {
		ui.Header("sync")
		err := jobs.Sync()
		if err != nil {
			ui.Die(err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(syncCmd)
}
