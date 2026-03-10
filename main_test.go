package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHelloHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	helloHandler(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status 200, got %d", resp.StatusCode)
	}

	contentType := resp.Header.Get("Content-Type")
	if !strings.Contains(contentType, "text/html") {
		t.Errorf("expected Content-Type to contain text/html, got %s", contentType)
	}

	body := w.Body.String()
	if !strings.Contains(body, message) {
		t.Errorf("expected body to contain %q", message)
	}

	if !strings.Contains(body, "<h1>") {
		t.Error("expected body to contain an h1 element")
	}
}
