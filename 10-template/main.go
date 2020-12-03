package main

import (
	"os"
	"text/template"
)

type User struct {
	FirstName string
	Age       uint
}

// 1- implement the stringer interface of fmt for User.

func main() {
	// 2- define a variable name u of type User with its fields.
	tmpl, err := template.New("test").Parse("Hello my name is {{.FirstName}} and I'm {{.Age}} years old.")
	if err != nil {
		panic(err)
	}
	// 3- replace the expression nil with u.
	err = tmpl.Execute(os.Stdout, nil)
	if err != nil {
		panic(err)
	}
}
