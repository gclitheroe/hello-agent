package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

const message = "Hello Atlassian from the Codex SDK"

var page = template.Must(template.New("home").Parse(`<!doctype html>
<html lang="en">
<head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <title>Hello Atlassian</title>
  <style>
    body {
      min-height: 100vh;
      margin: 0;
      display: grid;
      place-items: center;
      font-family: Arial, Helvetica, sans-serif;
      background: #ffffff;
    }

    h1 {
      margin: 0;
      color: #0052cc;
      font-size: 4rem;
      line-height: 1.1;
      text-align: center;
    }
  </style>
</head>
<body>
  <h1>{{.Message}}</h1>
</body>
</html>
`))

type pageData struct {
	Message string
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := page.Execute(w, pageData{Message: message}); err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", homeHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	addr := ":" + port
	fmt.Printf("Listening on http://localhost%s\n", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}
