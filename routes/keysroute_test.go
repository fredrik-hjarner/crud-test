package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/fredrik-hjarner/ztorage/diskv"
)

func setupFixture() {
	diskv.Diskv.EraseAll()
}

func TestEmptyKeysRoute(t *testing.T) {
	setupFixture()
	request, _ := http.NewRequest("GET", "/keys", nil)
	response := httptest.NewRecorder()
	HTTPHandler.ServeHTTP(response, request)
	AssertResponseCode(t, http.StatusOK, response.Code)
	// response.Body
}
