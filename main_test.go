package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	handler(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status 200, got %d", resp.StatusCode)
	}
	if ct := resp.Header.Get("Content-Type"); ct != "text/html" {
		t.Errorf("expected Content-Type text/html, got %s", ct)
	}
	body := w.Body.String()
	if !strings.Contains(body, "Hello Atlassian from Github Copilot") {
		t.Error("response body does not contain expected message")
	}
	if !strings.Contains(body, "<h1") {
		t.Error("response body does not contain a heading element for large font")
	}
}
