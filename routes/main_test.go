package routes

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/fredrik-hjarner/ztorage/diskv"
)

var HTTPHandler http.Handler

func TestMain(m *testing.M) {
	diskv.Diskv.EraseAll()
	HTTPHandler = CreateHttpHandler()
	code := m.Run()
	os.Exit(code)
}

func SetupFixture() {
	diskv.Diskv.EraseAll()
}

func SendSetValue(key string, value string) *httptest.ResponseRecorder {
	request, _ := http.NewRequest("POST", fmt.Sprintf("/value?key=%v&value=%v", key, value), nil)
	response := httptest.NewRecorder()
	HTTPHandler.ServeHTTP(response, request)
	return response
}
