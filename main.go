package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	html := `<!DOCTYPE html>
<html>
<head>
	<title>Hello EngFest</title>
	<style>
		body {
			display: flex;
			justify-content: center;
			align-items: center;
			min-height: 100vh;
			margin: 0;
			font-family: Arial, sans-serif;
			background-color: #f0f0f0;
		}
		.message {
			font-size: 72px;
			font-weight: bold;
			color: #333;
			text-align: center;
			padding: 20px;
		}
	</style>
</head>
<body>
	<div class="message">Hello EngFest #5 from Github Copilot</div>
</body>
</html>`
	fmt.Fprint(w, html)
}

func main() {
	http.HandleFunc("/", handler)
	port := ":8080"
	fmt.Printf("Server starting on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
