package main

import (
	"html/template"
	"log"
	"net/http"
)

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
            font-family: Arial, sans-serif;
            background-color: #f5f5f5;
        }
        .message {
            font-size: 3rem;
            font-weight: bold;
            color: #0052CC;
            text-align: center;
            padding: 2rem;
        }
    </style>
</head>
<body>
    <div class="message">Hello Atlassian from Github Copilot</div>
</body>
</html>`

var tmpl = template.Must(template.New("index").Parse(htmlTemplate))

func handleIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := tmpl.Execute(w, nil); err != nil {
		log.Printf("Error executing template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", handleIndex)
	
	port := ":8080"
	log.Printf("Starting server on http://localhost%s", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
