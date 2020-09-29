package requestmodels

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/asaskevich/govalidator"
)

func TestValid(t *testing.T) {
	createUserRequest := CreateUserRequest{
		FirstName: "FirstName",
		LastName:  "LastName",
		Email:     "hej@hej.com",
	}

	_, err := govalidator.ValidateStruct(createUserRequest)

	require.Equal(t, nil, err)
}

func TestTooShortFirstName(t *testing.T) {
	createUserRequest := CreateUserRequest{
		FirstName: "",
		LastName:  "LastName",
		Email:     "hej@hej.com",
	}

	_, err := govalidator.ValidateStruct(createUserRequest)

	require.NotEqual(t, nil, err)
}

func TestNumbersInFirstName(t *testing.T) {
	createUserRequest := CreateUserRequest{
		FirstName: "666",
		LastName:  "LastName",
		Email:     "hej@hej.com",
	}

	_, err := govalidator.ValidateStruct(createUserRequest)

	require.NotEqual(t, nil, err)
}

func TestIncorrectEmail(t *testing.T) {
	createUserRequest := CreateUserRequest{
		FirstName: "FirstName",
		LastName:  "LastName",
		Email:     "hej@hej",
	}

	_, err := govalidator.ValidateStruct(createUserRequest)

	require.NotEqual(t, nil, err)
}
