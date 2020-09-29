package models

// User model
type User struct {
	ID        string `json:"id" valid:"required,stringlength(1|255)"`
	FirstName string `json:"firstName" valid:"required,alpha,stringlength(1|255)"`
	LastName  string `json:"lastName" valid:"required,alpha,stringlength(1|255)"`
	Email     string `json:"email" valid:"required,email"`
}
