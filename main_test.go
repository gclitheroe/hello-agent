package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func get(t *testing.T, method, path string) *http.Response {
	t.Helper()

	rec := httptest.NewRecorder()
	newMux().ServeHTTP(rec, httptest.NewRequest(method, path, nil))
	return rec.Result()
}

func TestGreetServesTheMessage(t *testing.T) {
	res := get(t, http.MethodGet, "/")
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("status = %d, want %d", res.StatusCode, http.StatusOK)
	}

	if got, want := res.Header.Get("Content-Type"), "text/html; charset=utf-8"; got != want {
		t.Errorf("Content-Type = %q, want %q", got, want)
	}

	body := bodyOf(t, res)
	if !strings.Contains(body, message) {
		t.Errorf("body does not contain %q:\n%s", message, body)
	}
	if want := "Hello Jira I'm using skills"; message != want {
		t.Errorf("message = %q, want %q", message, want)
	}
}

// The work item asks for a large blue font centred on the page, so the styling
// is part of the requirement rather than incidental presentation.
func TestGreetIsLargeBlueAndCentred(t *testing.T) {
	res := get(t, http.MethodGet, "/")
	defer res.Body.Close()

	body := bodyOf(t, res)

	for _, style := range []string{
		"color: #0052cc",          // blue
		"font-size: clamp(",       // large, scaled to the viewport
		"align-items: center",     // centred vertically
		"justify-content: center", // centred horizontally
	} {
		if !strings.Contains(body, style) {
			t.Errorf("body is missing %q:\n%s", style, body)
		}
	}

	// The message must be the page heading, not buried in body text.
	if !strings.Contains(body, "<h1>"+message+"</h1>") {
		t.Errorf("message is not the <h1> heading:\n%s", body)
	}
}

func TestUnknownPathIsNotFound(t *testing.T) {
	res := get(t, http.MethodGet, "/nope")
	defer res.Body.Close()

	if res.StatusCode != http.StatusNotFound {
		t.Errorf("status = %d, want %d", res.StatusCode, http.StatusNotFound)
	}
}

func TestNonGetIsRejected(t *testing.T) {
	res := get(t, http.MethodPost, "/")
	defer res.Body.Close()

	if res.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("status = %d, want %d", res.StatusCode, http.StatusMethodNotAllowed)
	}
}

func bodyOf(t *testing.T, res *http.Response) string {
	t.Helper()

	b, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("reading body: %v", err)
	}
	return string(b)
}
