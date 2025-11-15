package model

import "time"

type User struct {
	FirstName string
	LastName  string
	Birthdate time.Time
	IsActive  bool
}

func (u *User) Modified(user *User) {
	user.FirstName = "Morteza"
}
func NewUser(firstName string, lastName string, birthdate time.Time, isActive bool) User {
	return User{
		FirstName: firstName,
		LastName:  lastName,
		Birthdate: birthdate,
		IsActive:  isActive,
	}
}
func Deactive(user *User) {
	user.IsActive = false
}
