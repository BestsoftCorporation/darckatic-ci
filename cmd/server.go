package cmd

import (
	"github.com/spf13/cobra"
)

var deployCmd = &cobra.Command{
	Use:     "deploy",
	Aliases: []string{"deploy"},
	Short:   "Deploy repository",
	Args:    cobra.ExactArgs(4),
	Run: func(cmd *cobra.Command, args []string) {

		println("Repository deployed successfully")
	},
}

func init() {
	rootCmd.AddCommand(listProjectsCmd)
	rootCmd.AddCommand(downloadCmd)
	rootCmd.AddCommand(listBranchesCmd)
}
