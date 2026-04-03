package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	html := `<!DOCTYPE html>
<html>
<head>
    <title>Hello Atlassian</title>
</head>
<body>
    <h1 style="color: blue; font-size: 48px;">Hello Atlassian from Claude Managed Agents</h1>
</body>
</html>`
	fmt.Fprint(w, html)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server starting on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
