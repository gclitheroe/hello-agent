package main

import (
	"fmt"
	"net/http"
)

const message = "Hello Atlassian from Github Copilot"

const htmlTemplate = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Hello Atlassian</title>
    <style>
        body {
            display: flex;
            justify-content: center;
            align-items: center;
            min-height: 100vh;
            margin: 0;
            background-color: #f5f5f5;
            font-family: Arial, sans-serif;
        }
        h1 {
            font-size: 4rem;
            color: #0052cc;
            text-align: center;
        }
    </style>
</head>
<body>
    <h1>%s</h1>
</body>
</html>`

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, htmlTemplate, message)
}

func main() {
	http.HandleFunc("/", helloHandler)
	fmt.Println("Server starting on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}
