package main

import (
	"fmt"
	"log"
	"net/http"
)

const message = "Hello Atlassian from Github Copilot"

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, `<!DOCTYPE html>
<html>
<head><title>Hello</title></head>
<body>
<h1 style="font-size:3em;">%s</h1>
</body>
</html>`, message)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
