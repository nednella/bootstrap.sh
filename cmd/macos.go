package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var macosCmd = &cobra.Command{
	Use:   "macos",
	Short: "Apply macOS preferences",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("macos: not implemented")
	},
}

func init() {
	rootCmd.AddCommand(macosCmd)
}
