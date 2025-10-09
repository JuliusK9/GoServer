package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	server := NewProcessServer(NewProcessStore())

	fmt.Println("Starting process manager server on :8080")
	log.Fatal(http.ListenAndServe(":8080", server))
}
