package cmd

import (
	"github.com/nednella/bootstrap.sh/internal/ui"
	"github.com/nednella/bootstrap.sh/internal/utils"
	"github.com/spf13/cobra"
)

var dryRun bool

var rootCmd = &cobra.Command{
	Use:   "bootstrap",
	Short: "Bootstrap a fresh macOS install",
	Long: `Bootstrap a fresh Mac quicker than the time it takes to make a cuppa.

Installs Homebrew + Brewfile, symlinks dotfiles into $HOME, and applies macOS preferences.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		utils.DryRun = dryRun
		ui.Banner()
		if dryRun {
			ui.Warn("DRY RUN — no changes will be made\n")
		}
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(
		&dryRun, "dry-run", "d", false,
		"preview actions without executing them",
	)
}
