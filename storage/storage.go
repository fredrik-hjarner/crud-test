package storage

import (
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
	Users = append(Users, user)
}
