package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start [name]",
	Short: "Start a new process",
	Long:  `Start a new process with a given name`,
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		serverURL := "http://localhost:8080"

		url := fmt.Sprintf("%s/process/%s", serverURL, name)
		req, err := http.NewRequest("POST", url, nil)
		if err != nil {
			log.Printf("Error creating request: %v", err)
			return
		}

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			log.Printf("Error calling server: %v", err)
			return
		}
		defer resp.Body.Close()

		switch resp.StatusCode {
		case http.StatusAccepted:
			fmt.Printf("Process '%s' started successfully\n", name)
		case http.StatusConflict:
			log.Printf("Process '%s' already exists", name)
		case http.StatusBadRequest:
			log.Printf("Invalid request: name parameter is required")
		default:
			log.Printf("Server returned error: %s", resp.Status)
		}
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
