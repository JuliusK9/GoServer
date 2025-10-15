package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	process "github.com/JuliusK9/GoServer"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ProcessManager",
	Short: "A CLI tool to manage processes",
	Long: `ProcessManager is a CLI tool that allows you to:
- Start new processes
- List running processes`,
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start the HTTP server",
	Run: func(cmd *cobra.Command, args []string) {
		store := process.NewProcessStore()
		server := process.NewProcessServer(store)

		fmt.Println("Starting server on :8080")
		log.Fatal(http.ListenAndServe(":8080", server))
	},
}

func main() {
	if len(os.Args) == 1 {
		os.Args = append(os.Args, "help")
	}
	rootCmd.AddCommand(startCmd, listCmd, serverCmd)
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
