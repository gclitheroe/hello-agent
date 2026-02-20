package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, rr.Code)
	}

	expected := "Hello Atlassian from Github Copilot"
	if !strings.Contains(rr.Body.String(), expected) {
		t.Errorf("expected body to contain %q, got %q", expected, rr.Body.String())
	}
}
