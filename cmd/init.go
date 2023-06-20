package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new Jif workspace",
	Long:  `Initialize a new Jif workspace`,
	Run: func(cmd *cobra.Command, args []string) {
		writeConfiguration()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}

type Config struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Team     string `json:"team"`
	Sprint   string `json:"sprint"`
}

func writeConfiguration() {
	fmt.Println("Writing configuration...")

	file, _ := json.MarshalIndent("", "", " ")
	_ = ioutil.WriteFile("config.json", file, 0644)
}
