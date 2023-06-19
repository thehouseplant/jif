package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "jif",
	Short: "An opinionated CLI tool for interacting with the Jira API",
	Long: "An opinionated CLI tool for interacting with the Jira API",
	RunE: func(cmd &cobra.Command, args []string) error {
		return cmd.Help()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
