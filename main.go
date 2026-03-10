package main

import (
	"fmt"
	"net/http"
)

const message = "Hello Atlassian from Github Copilot"

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Hello</title>
    <style>
        body {
            display: flex;
            justify-content: center;
            align-items: center;
            min-height: 100vh;
            margin: 0;
            font-family: sans-serif;
            background-color: #f0f4f8;
        }
        h1 {
            font-size: 4rem;
            color: #172b4d;
        }
    </style>
</head>
<body>
    <h1>%s</h1>
</body>
</html>`, message)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server listening on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
