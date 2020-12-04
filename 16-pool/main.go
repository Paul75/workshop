package main

import (
	"math/rand"
	"strings"
	"sync"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	var data sync.Pool

	// declare a variable name u of type User.
	// create a loop of 10e5 iterations.
	// in the for loop initiate the two fields of the variable u
	// the FirstName and the LastName should call the randomString function to get conent.
	// then Put this user into the data Pool.

	// declare a variable name u2 of type User.
	// create a loop of 10e5 iterations.
	// then Get the value form the data Pool.
	// if this value is not nil
	// convert this value from an interface into a User type
	// display this value with the fmt package and format the result with %#v
}

type User struct {
	FirstName string
	LastName  string
}

const all = "abcdefghijklmnopqrstuvwxyz"

func randomString() string {

	length := rand.Intn(15) + 2
	buf := make([]byte, length)

	for i := 0; i < length; i++ {
		buf[i] = all[rand.Intn(len(all))]
	}
	rand.Shuffle(len(buf), func(i, j int) {
		buf[i], buf[j] = buf[j], buf[i]
	})
	return strings.Title(string(buf))
}
