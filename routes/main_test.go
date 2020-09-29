package routes

import (
	"net/http"
	"os"
	"testing"
)

// HTTPHandler ...
var HTTPHandler http.Handler

// TestMain ...
func TestMain(m *testing.M) {
	HTTPHandler = CreateHTTPHandler()
	code := m.Run()
	os.Exit(code)
}

// SetupFixture ...
func SetupFixture() {
}
