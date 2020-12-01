package main

import "fmt"

func main() {
	// 1- fix the delared map and change from the key b the value into B
	// display the current map.
	fmt.Println(ma)
	// 2- create a map called mu of map of string of a User
	mu := make(map[string]User)
	// create a key called e15e98dc-33ca-11eb-b206-1f97348a99b3 and insert a User
	mu["e15e98dc-33ca-11eb-b206-1f97348a99b3"] = User{
		FirstName: "Rob",
		LastName:  "Pike",
	}
	// display the result.
	fmt.Println(mu)
	// 3- delete the current created key
	delete(mu, "e15e98dc-33ca-11eb-b206-1f97348a99b3")
	// and display if there is something into this current key.
	_, ok := mu["e15e98dc-33ca-11eb-b206-1f97348a99b3"]
	fmt.Println(ok)
}

var ma map[string]string = map[string]string{
	"a": "A",
	"b": "B",
	"c": "C",
	"d": "D",
}

type User struct {
	FirstName string
	LastName  string
}
