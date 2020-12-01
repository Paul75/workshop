package main

func main() {

}

type User struct {
	FirstName string
	LastName  string
}

// 1- create the constructor of User
func NewUser(fn, ln string) *User {
	return &User{
		FirstName: fn,
		LastName:  ln,
	}
}

// 2- create a method attached to User to display "Hello "+ it's first name.
func (u *User) Hello() string {
	return "Hello " + u.FirstName
}

// 3- create a method that has a param fn in string and mutate the user FristName into it's method.
func (u *User) SetFirstName(fn string) {
	u.FirstName = fn
}
