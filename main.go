package main

import (
	"fmt"
	"net/http"
)

const message = "Hello Atlassian from Github Copilot"

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, message)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
