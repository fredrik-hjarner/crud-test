package routes

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

/*
	Helper functions
*/

func SendGetKeys() *httptest.ResponseRecorder {
	request, _ := http.NewRequest("GET", "/keys", nil)
	response := httptest.NewRecorder()
	HTTPHandler.ServeHTTP(response, request)
	return response
}

func SendGetPrefixedKeys(prefix string) *httptest.ResponseRecorder {
	request, _ := http.NewRequest("GET", fmt.Sprintf("/keys?prefix=%v", prefix), nil)
	response := httptest.NewRecorder()
	HTTPHandler.ServeHTTP(response, request)
	return response
}

/*
	Tests
*/

func TestEmptyKeysRoute(t *testing.T) {
	SetupFixture()

	response := SendGetKeys()
	AssertResponseCode(t, http.StatusOK, response.Code)
	require.EqualValues(t, "[]", response.Body.String())
}

func TestAddKeyAndSeeIfItExists(t *testing.T) {
	SetupFixture()
	SendSetValue("somekey", "somevalue")

	response := SendGetKeys()
	AssertResponseCode(t, http.StatusOK, response.Code)
	require.EqualValues(t, `["somekey"]`, response.Body.String())
}

func TestGetPrefixedKeys(t *testing.T) {
	SetupFixture()
	SendSetValue("somekey", "somevalue")
	SendSetValue("anotherkey", "anothervalue")

	{
		response := SendGetPrefixedKeys("some")
		AssertResponseCode(t, http.StatusOK, response.Code)
		require.EqualValues(t, `["somekey"]`, response.Body.String())
	}
	{
		response := SendGetPrefixedKeys("another")
		AssertResponseCode(t, http.StatusOK, response.Code)
		require.EqualValues(t, `["anotherkey"]`, response.Body.String())
	}
}
