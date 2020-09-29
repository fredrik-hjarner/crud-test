package requestmodels

import "github.com/fredrik-hjarner/crud-test/models"

// UpdateUserRequest ...
type UpdateUserRequest struct {
	FirstName string `json:"firstName" valid:"required,alpha,stringlength(1|255)"`
	LastName  string `json:"lastName" valid:"required,alpha,stringlength(1|255)"`
	Email     string `json:"email" valid:"required,email"`
}

// ToUser ...
func (user *UpdateUserRequest) ToUser(ID string) models.User {
	return models.User{
		ID:        ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	}
}
