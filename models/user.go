package models

// User model
type User struct {
	ID        string `json:"id" valid:"required"`
	FirstName string `json:"firstName" valid:"required,alpha"`
	LastName  string `json:"lastName" valid:"required,alpha"`
	Email     string `json:"email" valid:"required,email"`
}
