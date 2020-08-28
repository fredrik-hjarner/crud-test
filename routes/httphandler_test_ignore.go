package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNonexistingUrl(t *testing.T) {
	request, _ := http.NewRequest("GET", "/nonexistingroute", nil)
	response := httptest.NewRecorder()
	HTTPHandler.ServeHTTP(response, request)
	AssertResponseCode(t, http.StatusNotFound, response.Code)
}

func TestExistingUrl(t *testing.T) {
	request, _ := http.NewRequest("GET", "/", nil)
	response := httptest.NewRecorder()
	HTTPHandler.ServeHTTP(response, request)
	AssertResponseCode(t, http.StatusOK, response.Code)
}
