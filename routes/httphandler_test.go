package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestNonexistingURL ...
func TestNonexistingURL(t *testing.T) {
	request, _ := http.NewRequest("GET", "/nonexistingroute", nil)
	response := httptest.NewRecorder()
	HTTPHandler.ServeHTTP(response, request)
	AssertResponseCode(t, http.StatusNotFound, response.Code)
}

// TestExistingURL ...
func TestExistingURL(t *testing.T) {
	request, _ := http.NewRequest("GET", "/", nil)
	response := httptest.NewRecorder()
	HTTPHandler.ServeHTTP(response, request)
	AssertResponseCode(t, http.StatusOK, response.Code)
}
