package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHelloHandlerRendersMessageAndBlueFont(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	helloHandler(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, rec.Code)
	}

	body := rec.Body.String()
	for _, want := range []string{
		"Hello Atlassian from the Codex SDK",
		"color: #0052cc",
		"font-size: clamp(3rem, 10vw, 8rem)",
	} {
		if !strings.Contains(body, want) {
			t.Fatalf("expected response body to contain %q", want)
		}
	}
}

func TestHelloHandlerNotFound(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/missing", nil)
	rec := httptest.NewRecorder()

	helloHandler(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Fatalf("expected status %d, got %d", http.StatusNotFound, rec.Code)
	}
}
