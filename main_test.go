package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHomeHandlerRendersMessageWithLargeBlueFont(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	homeHandler(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, rec.Code)
	}

	body := rec.Body.String()
	checks := []string{
		"Hello Atlassian from the Codex SDK",
		"color: #0052cc",
		"font-size: 4rem",
	}
	for _, want := range checks {
		if !strings.Contains(body, want) {
			t.Fatalf("expected response body to contain %q", want)
		}
	}

	contentType := rec.Header().Get("Content-Type")
	if !strings.HasPrefix(contentType, "text/html; charset=utf-8") {
		t.Fatalf("expected HTML content type, got %q", contentType)
	}
}

func TestHomeHandlerReturnsNotFoundForOtherPaths(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/missing", nil)
	rec := httptest.NewRecorder()

	homeHandler(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Fatalf("expected status %d, got %d", http.StatusNotFound, rec.Code)
	}
}
