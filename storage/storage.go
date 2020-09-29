package storage

import (
	"errors"

	"github.com/fredrik-hjarner/crud-test/models"
)

// Users A slice of all the Users in this (in-memory/non-persisted) storage
var Users = make([]models.User, 0)

var user = models.User{
	ID:        "1",
	FirstName: "Fredrik",
	LastName:  "Hjärner",
	Email:     "fredrik.hjarner@interactivesolutions.se",
}

// Init Initializes the storage
func Init() {
	// Users = append(Users, user)
}

// GetUserByID ...
func GetUserByID(id string) (models.User, error) {
	for _, user := range Users {
		if user.ID == id {
			return user, nil
		}
	}
	// TODO: should not return an empty User, should return nil, but then I have to use pointers.
	return models.User{}, errors.New("no user with that id")
}

// AddUser ...
func AddUser(user models.User) {
	Users = append(Users, user)
}

// ReplaceUser ...
func ReplaceUser(id string, newUser models.User) {
	indexOfUser := 0 // TODO: This sucks
	for i := 0; i < len(Users); i++ {
		if Users[i].ID == id {
			indexOfUser = i
		}
	}
	Users[indexOfUser] = newUser
}
