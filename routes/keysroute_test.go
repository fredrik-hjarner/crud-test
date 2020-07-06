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

func SendGetKeys(namespace string) *httptest.ResponseRecorder {
	request, _ := http.NewRequest("GET", fmt.Sprintf("/keys?namespace=%v", namespace), nil)
	response := httptest.NewRecorder()
	HTTPHandler.ServeHTTP(response, request)
	return response
}

func SendGetPrefixedKeys(namespace string, prefix string) *httptest.ResponseRecorder {
	request, _ := http.NewRequest("GET", fmt.Sprintf("/keys?namespace=%v&prefix=%v", namespace, prefix), nil)
	response := httptest.NewRecorder()
	HTTPHandler.ServeHTTP(response, request)
	return response
}

/*
	Tests
*/

func TestEmptyKeysRoute(t *testing.T) {
	SetupFixture()

	response := SendGetKeys("app")
	AssertResponseCode(t, http.StatusOK, response.Code)
	require.EqualValues(t, "[]", response.Body.String())
}

func TestAddKeyAndSeeIfItExists(t *testing.T) {
	SetupFixture()
	SendSetValue("app", "somekey", "somevalue")

	{
		response := SendGetKeys("app")
		AssertResponseCode(t, http.StatusOK, response.Code)
		require.EqualValues(t, `["somekey"]`, response.Body.String())
	}
	{
		response := SendGetKeys("someotherapp")
		AssertResponseCode(t, http.StatusOK, response.Code)
		require.EqualValues(t, `[]`, response.Body.String())
	}
}

func TestGetPrefixedKeys(t *testing.T) {
	SetupFixture()
	SendSetValue("app", "somekey", "somevalue")
	SendSetValue("app", "anotherkey", "anothervalue")

	{
		response := SendGetPrefixedKeys("someotherapp", "some")
		AssertResponseCode(t, http.StatusOK, response.Code)
		require.EqualValues(t, `[]`, response.Body.String())
	}
	{
		response := SendGetPrefixedKeys("app", "some")
		AssertResponseCode(t, http.StatusOK, response.Code)
		require.EqualValues(t, `["somekey"]`, response.Body.String())
	}
	{
		response := SendGetPrefixedKeys("app", "another")
		AssertResponseCode(t, http.StatusOK, response.Code)
		require.EqualValues(t, `["anotherkey"]`, response.Body.String())
	}
}
