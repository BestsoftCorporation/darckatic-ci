package cmd

import (
	parse "darkatic-ci/internal/config"
	"fmt"
	"github.com/spf13/cobra"
)

var reverseCmd = &cobra.Command{
	Use:     "init",
	Aliases: []string{"init"},
	Short:   "Init config file",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		result, err := parse.Init(args[0])
		if err != nil {
			fmt.Println("An error occuered: " + err.Error())
			return
		}

		println(result.Environment.Name)
	},
}

func init() {
	rootCmd.AddCommand(reverseCmd)
}
