package main

func main() {
	// 1- fix the delared map and change from the key b the value into B
	// display the current map.
	// 2- create a map called mu of map of string of a User
	// create a key called e15e98dc-33ca-11eb-b206-1f97348a99b3 and insert a User
	// display the result.
	// 3- delete the current created key
	// and display if there is something into this current key.
}

var ma map[string]string = map[string]string{
	"a": "A", "b": "Z", "c": "C", "d": "D",
}

type User struct {
	FirstName string
	LastName  string
}
