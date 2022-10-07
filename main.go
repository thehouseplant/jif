package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	jira "gopkg.in/andygrunwald/go-jira.v1"
)

func main() {
	// Define the self command
	var cmdSelf = &cobra.Command{
		Use:   "self",
		Short: "Get information on your Jira user",
		Long:  `This is a great way to ensure that you are authenticating properly`,
		Run: func(cmd *cobra.Command, args []string) {
			// Load local .env file
			godotenv.Load(".env")

			// Create BasicAuth transport object
			tp := jira.BasicAuthTransport{
				Username: os.Getenv("JIRA_USER"),
				Password: os.Getenv("JIRA_TOKEN"),
			}

			// Create new Jira Client
			client, err := jira.NewClient(tp.Client(), os.Getenv("JIRA_URL"))
			if err != nil {
				log.Fatal(err)
			}

			// Execute the GetSelf command
			me, _, err := client.User.GetSelf()
			if err != nil {
				log.Fatal(err)
			}

			// Log the results to the console
			fmt.Println(me)
		},
	}

	// Define root command and its children
	var rootCmd = &cobra.Command{Use: "jif"}
	rootCmd.AddCommand(cmdSelf)
	rootCmd.Execute()
}
