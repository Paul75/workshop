package main

import "fmt"

func main() {
	fmt.Println(tbl1)
	fmt.Println(tbl2)

	// 3- change the value at index 0 to a and display the result.
	fmt.Println(tbl2)

	// 4- copy all values from tbl2 into tbl1 and display the values.
	fmt.Println(tbl1)
}

// 1- fix this bug so the array has 4 elements.
var tbl1 [8]string = [4]string{
	"1", "2", "3", "4",
}

// 2- fix the bug so the array contain string.
var tbl2 [8]int = [8]int{
	"z", "b", "c", "d",
}
