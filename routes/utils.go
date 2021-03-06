package routes

import (
	"testing"
)

// AssertResponseCode ...
func AssertResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}
