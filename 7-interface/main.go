package main

import "fmt"

func main() {
	// implement the interface contract Changer by adding all methods set form it into User.
	// the functions SetFirstName and SetLastName should respectivly mutate the user firstName and lastName.
	// initiate a user with the fields
	u := User{
		firstName: "Rob",
		lastName:  "Pike",
	}
	// user one of the function and display the reslut.
	u.SetFirstName("Robert")
	fmt.Println(u)
}

var (
	_ Changer      = &User{}
	_ fmt.Stringer = &User{}
)

type User struct {
	firstName string
	lastName  string
}

func (u *User) String() string {
	return u.firstName
}

func (u *User) SetFirstName(fn string) {
	u.firstName = fn
}

func (u *User) SetLastName(ln string) {
	u.lastName = ln
}

type Changer interface {
	SetFirstName(fn string)
	SetLastName(ln string)
}
