// Command hello-agent serves a single page displaying "Hello Atlassian".
package main

import (
	"flag"
	"log"
	"net/http"
)

const page = `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="utf-8">
<meta name="viewport" content="width=device-width, initial-scale=1">
<title>Hello Atlassian</title>
<style>
  html, body {
    height: 100%;
    margin: 0;
  }
  body {
    display: flex;
    align-items: center;
    justify-content: center;
    background: #ffffff;
    font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, Helvetica, Arial, sans-serif;
  }
  h1 {
    margin: 0;
    color: #0052CC;
    font-size: 6rem;
    font-weight: 700;
    letter-spacing: -0.02em;
    text-align: center;
  }
  @media (max-width: 700px) {
    h1 { font-size: 3rem; }
  }
</style>
</head>
<body>
  <h1>Hello Atlassian</h1>
</body>
</html>
`

// helloHandler writes the greeting page. Any path other than "/" is a 404.
func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	if r.Method != http.MethodGet && r.Method != http.MethodHead {
		w.Header().Set("Allow", "GET, HEAD")
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if _, err := w.Write([]byte(page)); err != nil {
		log.Printf("writing response: %v", err)
	}
}

func newMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", helloHandler)
	return mux
}

func main() {
	addr := flag.String("addr", ":8080", "address to listen on")
	flag.Parse()

	log.Printf("listening on %s", *addr)
	if err := http.ListenAndServe(*addr, newMux()); err != nil {
		log.Fatal(err)
	}
}
