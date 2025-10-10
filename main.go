package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	log.Fatal(RunServer(":8080"))
}

func RunServer(addr string) error {
	server := NewProcessServer(NewProcessStore())
	fmt.Println("Starting process manager server on ", addr)
	return http.ListenAndServe(addr, server)
}
