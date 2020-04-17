package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETHealth(t *testing.T) {
	request, _ := http.NewRequest(http.MethodGet, "/health/", nil)
	response := httptest.NewRecorder()

	healthHandler(response, request)

	got := response.Body.String()
	want := "hello"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
