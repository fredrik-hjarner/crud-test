package routes

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var httpHandler http.Handler

func TestMain(m *testing.M) {
	httpHandler = CreateHttpHandler()
	code := m.Run()
	os.Exit(code)
}

func TestNonexistingUrl(t *testing.T) {
	request, _ := http.NewRequest("GET", "/nonexistingroute", nil)
	response := httptest.NewRecorder()
	httpHandler.ServeHTTP(response, request)
	AssertResponseCode(t, http.StatusNotFound, response.Code)
}

func TestExistingUrl(t *testing.T) {
	request, _ := http.NewRequest("GET", "/", nil)
	response := httptest.NewRecorder()
	httpHandler.ServeHTTP(response, request)
	AssertResponseCode(t, http.StatusOK, response.Code)
}
