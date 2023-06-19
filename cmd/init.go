package cmd

import (
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new Jif workspace",
	Long:  `Initialize a new Jif workspace`,
	Run: func(cmd *cobra.Command, args []string) {
		//
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
