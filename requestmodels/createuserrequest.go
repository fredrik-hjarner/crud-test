package requestmodels

import (
	"github.com/fredrik-hjarner/crud-test/models"
	"github.com/fredrik-hjarner/crud-test/utils"
)

// CreateUserRequest ...
type CreateUserRequest struct {
	FirstName string `json:"firstName" valid:"required,alpha"`
	LastName  string `json:"lastName" valid:"required,alpha"`
	Email     string `json:"email" valid:"required,email"`
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
