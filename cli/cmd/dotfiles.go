package cmd

import (
	"github.com/nednella/bootstrap.sh/internal/jobs"
	"github.com/nednella/bootstrap.sh/internal/ui"
	"github.com/spf13/cobra"
)

var dotfilesUndo bool

var dotfilesCmd = &cobra.Command{
	Use:     "dotfiles",
	Short:   "Symlink dotfiles into $HOME / XDG",
	GroupID: jobsGroupID,
	Run: func(cmd *cobra.Command, args []string) {
		ui.Header("dotfiles")
		var err error
		if dotfilesUndo {
			err = jobs.DotfilesUndo()
		} else {
			err = jobs.Dotfiles()
		}
		if err != nil {
			ui.Die(err.Error())
		}
	},
}

func init() {
	dotfilesCmd.Flags().BoolVarP(&dotfilesUndo, "undo", "u", false, "unlink dotfiles and restore the latest backup")
	rootCmd.AddCommand(dotfilesCmd)
}
