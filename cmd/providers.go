package cmd

import (
	provider "darkatic-ci/internal/provider"
	"fmt"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
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
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(downloadCmd)
}
