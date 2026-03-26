package main

import (
	"fmt"
	"log"
	"net/http"
)

const message = "Hello Atlassian from Github Copilot Coding Agent"

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Hello</title>
</head>
<body>
    <h1 style="font-size: 4em;">%s</h1>
</body>
</html>`, message)
}

func main() {
	http.HandleFunc("/", handler)
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
