package cmd

import (
	provider "darkatic-ci/internal/provider"
	"fmt"
	"github.com/spf13/cobra"
)

var listBranchesCmd = &cobra.Command{
	Use:     "list-branches",
	Aliases: []string{"list-branches"},
	Short:   "List branches",
	Args:    cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		hubProvider := provider.GitHubProvider{}
		result, err := hubProvider.ListBranches(args[0], args[1], args[2])
		if err != nil {
			fmt.Println("An error occuered: " + err.Error())
			return
		}

		for _, project := range result {
			fmt.Println(project)
		}
	},
}

var listProjectsCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"list"},
	Short:   "List repositories",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		hubProvider := provider.GitHubProvider{}
		result, err := hubProvider.FetchProjects(args[0], args[1])
		if err != nil {
			fmt.Println("An error occuered: " + err.Error())
			return
		}

		for _, project := range result {
			fmt.Println(project)
		}
	},
}

var downloadCmd = &cobra.Command{
	Use:     "download",
	Aliases: []string{"download"},
	Short:   "Download repository",
	Args:    cobra.ExactArgs(4),
	Run: func(cmd *cobra.Command, args []string) {
		hubProvider := provider.GitHubProvider{}
		err := hubProvider.DownloadZip(args[0], args[1], args[2], args[3])
		if err != nil {
			fmt.Println("An error occuered: " + err.Error())
			return
		}

		println("Project downloaded")
	},
}

func init() {
	rootCmd.AddCommand(listProjectsCmd)
	rootCmd.AddCommand(downloadCmd)
	rootCmd.AddCommand(listBranchesCmd)
}
