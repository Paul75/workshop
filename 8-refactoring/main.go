package main

import (
	"strings"

	"github.com/google/uuid"
)

func main() {
	// 1 - create a file called model into this current file
	// 2 - put all struct and methods declared here below into this new model file.
	// 3 - go back to this current file main.go
	// 4 - and create a new User called u with it's constructor into the main function.
	// 5 - use the function fmt.Println to display u.
	// 6 - use the method of u SayHello
	// 7 - create a Pet with it's constructor and name it p.
	// 8 - add this new pet to the last created User with the function AddPet.
	// 9 - use the function fmt.Println to display u.
}

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

func (u *User) AddPet(p Pet) {
	u.Pets = append(u.Pets, p)
}

func (u *User) String() string {
	if len(u.Pets) == 0 {
		return "Hello my name is " + u.FirstName
	}
	res := "Hi I'm" + u.FirstName + " and let me introduce you to "
	for _, p := range u.Pets {
		res += p.Name + ", "
	}
	return res[:len(res)-2]
}

type Pet struct {
	ID          string
	Name        string
	Description string
}

func NewPet(name, desc string) *Pet {
	return &Pet{
		ID:          uuid.New().String(),
		Name:        name,
		Description: desc,
	}
}
