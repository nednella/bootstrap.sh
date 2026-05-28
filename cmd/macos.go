package cmd

import (
	"fmt"

	"github.com/nednella/bootstrap.sh/internal/ui"
	"github.com/spf13/cobra"
)

var macosCmd = &cobra.Command{
	Use:   "macos",
	Short: "Apply macOS preferences",
	Run: func(cmd *cobra.Command, args []string) {
		ui.Header("macos")
		fmt.Println("macos: not implemented")
	},
}

func init() {
	rootCmd.AddCommand(macosCmd)
}
