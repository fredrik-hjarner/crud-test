package routes

import (
	"net/http"
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
