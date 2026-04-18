package main

import (
	"fmt"
	"log"
	"net/http"
)

const message = "Hello Atlassian from the Claude Agent SDK"

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, message)
}

func main() {
	http.HandleFunc("/", handler)
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
