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

func NewUser(firstName, lastName string) *User {
	return &User{
		FirstName: firstName,
		LastName:  lastName,
	}
}

func Deactive(user *User) {
	user.IsActive = false
}
