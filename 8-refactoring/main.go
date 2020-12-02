package main

import (
	"fmt"
	"workshop/8-refactoring/model"
)

func main() {
	// 4 - and create a new User called u with it's constructor into the main function.
	u := model.NewUser("Bob", "Pike", 50)
	// 5 - use the function fmt.Println to display u.
	fmt.Println(u)
	// 6 - use the method of u SayHello
	u.SayHello("Hi")
	// 7 - create a Pet with it's constructor and name it p.
	p := model.NewPet("Benji", "black dog")
	p2 := model.NewPet("Mistigris", "white cat")
	// 8 - add this new pet to the last created User with the function AddPet.
	u.AddPet(p)
	u.AddPet(p2)
	// 9 - use the function fmt.Println to display u.
	fmt.Println(u)
}
