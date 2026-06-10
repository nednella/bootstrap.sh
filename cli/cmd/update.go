package cmd

import (
	"github.com/nednella/bootstrap.sh/internal/jobs"
	"github.com/nednella/bootstrap.sh/internal/ui"
	"github.com/spf13/cobra"
)

var (
	updateTag  string
	updateList bool
)

var updateCmd = &cobra.Command{
	Use:     "update",
	Short:   "Update the binary to the latest release",
	GroupID: jobsGroupID,
	Run: func(cmd *cobra.Command, args []string) {
		ui.Header("update")
		var err error
		switch {
		case updateList:
			err = jobs.UpdateList()
		case updateTag != "":
			err = jobs.UpdateTag(updateTag)
		default:
			err = jobs.Update()
		}
		if err != nil {
			ui.Die(err.Error())
		}
	},
}

func init() {
	updateCmd.Flags().StringVarP(&updateTag, "tag", "t", "", "install a specific release")
	updateCmd.Flags().BoolVarP(&updateList, "list", "l", false, "list available releases")
	updateCmd.MarkFlagsMutuallyExclusive("tag", "list")
	rootCmd.AddCommand(updateCmd)
}
