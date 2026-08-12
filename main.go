// Command hello-agent serves a single page that greets the reader.
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

// message is the greeting shown in the centre of the page.
const message = "Hello Jira I'm using skills"

// pageTemplate holds the document served at /. The blue is Atlassian brand blue,
// and the font size scales with the viewport so the message stays large without
// overflowing a narrow window.
const pageTemplate = `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="utf-8">
<meta name="viewport" content="width=device-width, initial-scale=1">
<title>%s</title>
<style>
html,
body {
	height: 100%%;
	margin: 0;
}

body {
	display: flex;
	align-items: center;
	justify-content: center;
	background: #ffffff;
	font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Helvetica, Arial, sans-serif;
}

h1 {
	margin: 0;
	padding: 0 1rem;
	color: #0052cc;
	font-size: clamp(2.5rem, 9vw, 8rem);
	font-weight: 700;
	text-align: center;
}
</style>
</head>
<body>
<h1>%s</h1>
</body>
</html>
`

// page is the response body. The content never varies, so it is built once at
// start-up rather than rendered per request.
var page = fmt.Appendf(nil, pageTemplate, message, message)

func newMux() *http.ServeMux {
	mux := http.NewServeMux()
	// {$} matches / exactly, so any other path falls through to a 404.
	mux.HandleFunc("GET /{$}", greet)
	return mux
}

func greet(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if _, err := w.Write(page); err != nil {
		log.Printf("writing response: %v", err)
	}
}

func main() {
	addr := flag.String("addr", ":8080", "host:port to listen on")
	flag.Parse()

	log.Printf("listening on %s", *addr)
	if err := http.ListenAndServe(*addr, newMux()); err != nil {
		log.Fatal(err)
	}
}
