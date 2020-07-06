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

func SendGetValueByKey(key string) *httptest.ResponseRecorder {
	request, _ := http.NewRequest("GET", fmt.Sprintf("/value?key=%v", key), nil)
	response := httptest.NewRecorder()
	HTTPHandler.ServeHTTP(response, request)
	return response
}

func SendSetValue(key string, value string) *httptest.ResponseRecorder {
	request, _ := http.NewRequest("POST", fmt.Sprintf("/value?key=%v&value=%v", key, value), nil)
	response := httptest.NewRecorder()
	HTTPHandler.ServeHTTP(response, request)
	return response
}

func SendDeleteOneValue(key string) *httptest.ResponseRecorder {
	request, _ := http.NewRequest("DELETE", fmt.Sprintf("/value?key=%v", key), nil)
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

	response := SendSetValue("somekey", "somevalue")
	AssertResponseCode(t, http.StatusOK, response.Code)
}

func TestGetNonexistingValue(t *testing.T) {
	SetupFixture()

	response := SendGetValueByKey("somekey")
	AssertResponseCode(t, http.StatusNotFound, response.Code)
}

func TestGetExistingValue(t *testing.T) {
	SetupFixture()
	SendSetValue("somekey", "somevalue")

	response := SendGetValueByKey("somekey")
	AssertResponseCode(t, http.StatusOK, response.Code)
	require.EqualValues(t, "somevalue", response.Body.String())
}

func TestDeleteNonexistingValue(t *testing.T) {
	SetupFixture()

	response := SendDeleteOneValue("somekey")
	AssertResponseCode(t, http.StatusOK, response.Code)
}

func TestDeleteExistingValue(t *testing.T) {
	SetupFixture()
	SendSetValue("somekey", "somevalue")

	response := SendDeleteOneValue("somekey")
	AssertResponseCode(t, http.StatusOK, response.Code)
}

func TestDeleteAllValues(t *testing.T) {
	SetupFixture()
	SendSetValue("somekey", "somevalue")

	{
		response := SendDeleteAllValues()
		AssertResponseCode(t, http.StatusOK, response.Code)
	}
	{
		response := SendGetKeys()
		AssertResponseCode(t, http.StatusOK, response.Code)
		require.EqualValues(t, "[]", response.Body.String())
	}
}
