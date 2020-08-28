package requestmodels

import (
	"github.com/fredrik-hjarner/crud-test/models"
	"github.com/fredrik-hjarner/crud-test/utils"
)

// CreateUserRequest ...
type CreateUserRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

// ToUser ...
func (user *CreateUserRequest) ToUser() models.User {
	return models.User{
		ID:        utils.CreateUUID(),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	}
}
