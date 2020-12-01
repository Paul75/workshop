package main

import "fmt"

func main() {
	fmt.Println(tbl1)
	fmt.Println(tbl2)

	// 3- change the value at index 0 to a and display the result.
	tbl2[0] = "a"
	fmt.Println(tbl2)

	// 4- copy all values from tbl2 into tbl1 and display the values.
	for k, v := range tbl2 {
		if k > len(tbl1) {
			break
		}
		tbl1[k] = v
	}

	for i := 0; i < len(tbl1); i++ {
		tbl1[i] = tbl2[i]
	}
	fmt.Println(tbl1)
}

// 1- fix this bug so the array has 4 elements.
var tbl1 [4]string = [4]string{
	"1", "2", "3", "4",
}

// 2- fix the bug so the array contain string.
var tbl2 [8]string = [8]string{
	"z", "b", "c", "d",
}
