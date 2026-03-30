package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `<!DOCTYPE html>
<html>
<head><title>Hello EngFest</title></head>
<body>
<h1 style="font-size: 5em;">Hello EngFest #5 from Github Copilot</h1>
</body>
</html>`)
}

func main() {
	http.HandleFunc("/", handler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
