package cmd

import (
	"github.com/nednella/bootstrap.sh/internal/jobs"
	"github.com/nednella/bootstrap.sh/internal/ui"
	"github.com/spf13/cobra"
)

var dotfilesCmd = &cobra.Command{
	Use:     "dotfiles",
	Short:   "Symlink dotfiles into $HOME / XDG",
	GroupID: jobsGroupID,
	Run: func(cmd *cobra.Command, args []string) {
		ui.Header("dotfiles")
		err := jobs.Dotfiles()
		if err != nil {
			ui.Die(err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(dotfilesCmd)
}
