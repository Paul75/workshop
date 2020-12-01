package main

import "time"

func main() {
	// 1- compose User with DBFields
	// 2- compose User with the OptionalFields struct
	// 3- create a constructor of User with the optional fields.
}

type User struct {
	FirstName string
	LastName  string
	DBFields
	OptionalFields
}

func NewUser(fn, ln string, optional *OptionalFields) *User {
	var u User
	u.FirstName = fn
	u.LastName = ln
	if optional != nil {
		u.OptionalFields = *optional
	}
	return &u
}

type DBFields struct {
	ID         string
	CreatDate  time.Time
	UpdateDate time.Time
	DeleteDate time.Time
}

type OptionalFields struct {
	Email          string
	EmailValidated bool
	IsConnected    bool
}
