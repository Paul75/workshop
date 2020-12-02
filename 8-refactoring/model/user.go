package model

import (
	"errors"
	"strings"

	"github.com/google/uuid"
)

type User struct {
	ID        string
	FirstName string
	LastName  string
	Age       int
	Pets      []Pet
}

func NewUser(fn, ln string, age int) *User {
	return &User{
		ID:        uuid.New().String(),
		FirstName: fn,
		LastName:  ln,
		Age:       age,
	}
}

func (u *User) SayHello(msg string) string {
	return u.FirstName + " " + strings.ToUpper(u.LastName) + ": " + msg
}

func (u *User) AddPet(p *Pet) error {
	if p == nil {
		return errors.New("no pets to add")
	}
	u.Pets = append(u.Pets, *p)
	return nil
}

func (u *User) String() string {
	if len(u.Pets) == 0 {
		return "Hello my name is " + u.FirstName
	}
	res := "Hi I'm " + u.FirstName + " and let me introduce you to "
	for _, p := range u.Pets {
		res += p.Name + ", "
	}
	return res[:len(res)-2]
}
