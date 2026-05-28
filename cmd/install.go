package cmd

import (
	"github.com/nednella/bootstrap.sh/internal/jobs"
	"github.com/nednella/bootstrap.sh/internal/ui"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:     "install",
	Short:   "Install packages from Brewfile",
	GroupID: jobsGroupID,
	Run: func(cmd *cobra.Command, args []string) {
		ui.Header("install")
		err := jobs.Install()
		if err != nil {
			ui.Die(err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(installCmd)
}
