package cmd

import (
	"fmt"

	"github.com/nednella/bootstrap.sh/internal"
	"github.com/nednella/bootstrap.sh/internal/jobs"
	"github.com/nednella/bootstrap.sh/internal/ui"
	"github.com/nednella/bootstrap.sh/internal/utils"
	"github.com/spf13/cobra"
)

const jobsGroupID = "jobs"

var dryRun bool

var rootCmd = &cobra.Command{
	Use:   "bootstrap",
	Short: "Bootstrap a fresh macOS install",
	Long: `Bootstrap a fresh Mac quicker than the time it takes to make a cuppa.

Installs Homebrew + Brewfile, symlinks dotfiles into $HOME, and applies macOS preferences.`,
	Version: internal.Version,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		utils.DryRun = dryRun
		ui.Banner()
		if dryRun {
			fmt.Println()
			ui.Warn("DRY RUN — no changes will be made")
		}

		if cmd.GroupID != jobsGroupID {
			return
		}

		err := jobs.Preflight()
		if err != nil {
			ui.Die(err.Error())
		}
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddGroup(&cobra.Group{
		ID:    jobsGroupID,
		Title: "Job Commands",
	})
	rootCmd.PersistentFlags().BoolVarP(
		&dryRun, "dry-run", "d", false,
		"preview actions without executing them",
	)
}
