package models

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnmarshal(t *testing.T) {
	user := &User{}

	jsonData := `{
		"id": "1",
		"firstName": "Fredrik",
		"lastName": "Hjärner",
		"email": "fredrik.hjarner@interactivesolutions"
	}`

	require.NoError(t, json.Unmarshal([]byte(jsonData), user))

	require.Equal(t, user.ID, "1")
	require.Equal(t, user.FirstName, "Fredrik")
	require.Equal(t, user.LastName, "Hjärner")
	require.Equal(t, user.Email, "fredrik.hjarner@interactivesolutions")
}

func TestMarshal(t *testing.T) {
	user := &User{
		ID:        "1",
		FirstName: "Fredrik",
		LastName:  "Hjärner",
		Email:     "fredrik.hjarner@interactivesolutions",
	}

	expected := `{"id":"1","firstName":"Fredrik","lastName":"Hjärner","email":"fredrik.hjarner@interactivesolutions"}`

	result, err := json.Marshal(user)

	require.Equal(t, nil, err)
	require.Equal(t, expected, string(result))
}
