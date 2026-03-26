package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintln(w, `<!DOCTYPE html>
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
            font-family: sans-serif;
            background-color: #f5f5f5;
        }
        h1 {
            font-size: 4rem;
            color: #0052cc;
        }
    </style>
</head>
<body>
    <h1>Hello Atlassian from Github Copilot</h1>
</body>
</html>`)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Listening on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}
