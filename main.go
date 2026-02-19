package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Atlassian from Github Copilot")
}

func main() {
	http.HandleFunc("/", helloHandler)
	
	port := ":8080"
	log.Printf("Starting server on port %s", port)
	
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
