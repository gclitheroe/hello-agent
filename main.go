package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const message = "Hello Atlassian from the Codex SDK"

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", helloHandler)

	addr := ":" + port()
	log.Printf("listening on http://localhost%s", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal(err)
	}
}

func port() string {
	if value := os.Getenv("PORT"); value != "" {
		return value
	}
	return "8080"
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, `<!doctype html>
<html lang="en">
<head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <title>%[1]s</title>
  <style>
    html, body {
      height: 100%%;
    }

    body {
      margin: 0;
      display: grid;
      place-items: center;
      font-family: Arial, Helvetica, sans-serif;
      background: #ffffff;
    }

    h1 {
      margin: 0;
      color: #0052cc;
      font-size: clamp(3rem, 10vw, 8rem);
      font-weight: 800;
      text-align: center;
      line-height: 1.05;
    }
  </style>
</head>
<body>
  <h1>%[1]s</h1>
</body>
</html>
`, message)
}
