package main

import (
	"fmt"
	"log"
	"net/http"

	process "github.com/JuliusK9/GoServer"
)

func main() {
	store := process.NewProcessStore()
	server := process.NewProcessServer(store)

	fmt.Println("Starting web server on :8080")
	log.Fatal(http.ListenAndServe(":8080", server))
}
