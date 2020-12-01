package main

import "fmt"

func main() {
	// 3- display the values of tbl1
	fmt.Println(tbl1)
	// 4- copy all values from the tbl1 into a new slice called tbl2
	tbl2 := make([]string, len(tbl1))
	copy(tbl2, tbl1)
	// and display it's values of tbl2
	fmt.Println(tbl2)
	// change a value from tbl1 and display all values from tbl2
	tbl1[0] = "x"
	fmt.Println(tbl2)
	// 5- slice all values from the tbl1 into a new slice called tbl4
	tbl4 := tbl1[:]
	// and display it's values of tbl4
	fmt.Println(tbl4)
	// change a value from tbl1 and display all values from tbl4
	tbl1[1] = "w"
	fmt.Println(tbl4)
	// 6- create a slice of string with the following lenght 5 and capacity 7 called tbl5
	tbl5 := make([]string, 5, 7)
	// and display it's lenght and capacity.
	fmt.Println(len(tbl5), cap(tbl5))
}

// 1- fix the bug so the slice is valid in the declaration.
var tbl1 []string = []string{
	"z", "b", "c", "d",
}

// 2- fix the bug so the slice is valid in the declaration.
var tbl3 []string = []string{
	"z", "b", "c", "d",
}
