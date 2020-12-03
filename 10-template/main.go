package main

import (
	"os"
	"text/template"
)

type User struct {
	FirstName string
	Age       uint
}

func main() {
	// 1- define a variable name u of type User with its fields.
	tmpl, err := template.New("test").Parse("Hello my name is {{.FirstName}} and I'm {{.Age}} years old.")
	if err != nil {
		panic(err)
	}
	// 2- replace the expression nil with u.
	err = tmpl.Execute(os.Stdout, nil)
	if err != nil {
		panic(err)
	}
}
