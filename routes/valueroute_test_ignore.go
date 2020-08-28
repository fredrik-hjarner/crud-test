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

func SendGetValueByKey(namespace string, key string) *httptest.ResponseRecorder {
	request, _ := http.NewRequest("GET", fmt.Sprintf("/value?namespace=%v&key=%v", namespace, key), nil)
	response := httptest.NewRecorder()
	HTTPHandler.ServeHTTP(response, request)
	return response
}

func SendSetValue(namespace string, key string, value string) *httptest.ResponseRecorder {
	request, _ := http.NewRequest("POST", fmt.Sprintf("/value?namespace=%v&key=%v&value=%v", namespace, key, value), nil)
	response := httptest.NewRecorder()
	HTTPHandler.ServeHTTP(response, request)
	return response
}

func SendDeleteOneValue(namespace string, key string) *httptest.ResponseRecorder {
	request, _ := http.NewRequest("DELETE", fmt.Sprintf("/value?namespace=%v&key=%v", namespace, key), nil)
	response := httptest.NewRecorder()
	HTTPHandler.ServeHTTP(response, request)
	return response
}

func SendDeleteAllValues() *httptest.ResponseRecorder {
	request, _ := http.NewRequest("DELETE", "/value", nil)
	response := httptest.NewRecorder()
	HTTPHandler.ServeHTTP(response, request)
	return response
}

/*
	Tests
*/

func TestSetValue(t *testing.T) {
	SetupFixture()

	response := SendSetValue("app", "somekey", "somevalue")
	AssertResponseCode(t, http.StatusOK, response.Code)
}

func TestGetNonexistingValue(t *testing.T) {
	SetupFixture()

	response := SendGetValueByKey("app", "somekey")
	AssertResponseCode(t, http.StatusNotFound, response.Code)
}

func TestGetExistingValue(t *testing.T) {
	SetupFixture()
	SendSetValue("app", "somekey", "somevalue")

	{
		response := SendGetValueByKey("app", "somekey")
		AssertResponseCode(t, http.StatusOK, response.Code)
		require.EqualValues(t, "somevalue", response.Body.String())
	}
	{
		response := SendGetValueByKey("someotherapp", "somekey")
		AssertResponseCode(t, http.StatusNotFound, response.Code)
	}
}

func TestDeleteNonexistingValue(t *testing.T) {
	SetupFixture()

	response := SendDeleteOneValue("app", "somekey")
	AssertResponseCode(t, http.StatusOK, response.Code)
}

func TestDeleteExistingValue(t *testing.T) {
	SetupFixture()
	SendSetValue("app", "somekey", "somevalue")

	{
		response := SendGetKeys("app")
		AssertResponseCode(t, http.StatusOK, response.Code)
		require.EqualValues(t, `["somekey"]`, response.Body.String())
	}
	{
		response := SendDeleteOneValue("app", "somekey")
		AssertResponseCode(t, http.StatusOK, response.Code)
	}
	{
		response := SendGetKeys("app")
		AssertResponseCode(t, http.StatusOK, response.Code)
		require.EqualValues(t, `[]`, response.Body.String())
	}
}

func TestDeleteAllValues(t *testing.T) {
	SetupFixture()
	SendSetValue("app", "somekey", "somevalue")

	{
		response := SendDeleteAllValues()
		AssertResponseCode(t, http.StatusOK, response.Code)
	}
	{
		response := SendGetKeys("app")
		AssertResponseCode(t, http.StatusOK, response.Code)
		require.EqualValues(t, "[]", response.Body.String())
	}
}
