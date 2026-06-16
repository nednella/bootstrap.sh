package cmd

import (
	"github.com/nednella/bootstrap.sh/internal/jobs"
	"github.com/nednella/bootstrap.sh/internal/ui"
	"github.com/spf13/cobra"
)

var sshKeyCmd = &cobra.Command{
	Use:     "ssh-key",
	Short:   "Generate an SSH key",
	GroupID: jobsGroupID,
	Run: func(cmd *cobra.Command, args []string) {
		ui.Header("ssh-key")
		err := jobs.SSHKey()
		if err != nil {
			ui.Die(err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(sshKeyCmd)
}
