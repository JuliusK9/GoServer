package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all running processes",
	Long:  `List all currently running processes`,
	Run: func(cmd *cobra.Command, args []string) {
		serverURL := "http://localhost:8080"

		resp, err := http.Get(serverURL + "/process/list")
		if err != nil {
			log.Printf("Error calling server: %v", err)
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			log.Printf("Server returned error: %s", resp.Status)
			return
		}

		var processNames []string
		err = json.NewDecoder(resp.Body).Decode(&processNames)
		if err != nil {
			log.Printf("Error decoding response: %v", err)
			return
		}

		if len(processNames) == 0 {
			fmt.Println("No running processes")
			return
		}

		for _, name := range processNames {
			fmt.Printf("%s\n", name)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
