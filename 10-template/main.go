package main

import (
	"fmt"
	"os"
	"text/template"
)

type User struct {
	FirstName string
	Age       uint
}

func (u User) String() string {
	if u.Age < 10 {
		return fmt.Sprintf("Hi I'm %v and I have %v.", u.FirstName, u.Age)
	}
	return fmt.Sprintf("Hello my name is %v and I'm %v years old.", u.FirstName, u.Age)
}

func main() {
	// 1- define a variable name u of type User with its fields.
	u := User{
		FirstName: "Dennis",
		Age:       9,
	}
	data := map[string]interface{}{
		"User": u,
	}
	tmpl, err := template.New("test").Parse("{{.User}}")
	if err != nil {
		panic(err)
	}
	// 2- replace the expression nil with u.
	err = tmpl.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
}
