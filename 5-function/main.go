package main

func main() {}

// 1- create the function signature that has one param str in string and return a string.
func Hello(str string) string {
	return str
}

// 2- create the function signature that has a variadic params concatenate the result and return it.
func Concatenate(n ...string) string {
	res := ""
	for _, v := range n {
		res += v
	}
	return res
}

// 3- create a function signature that return another function that return an int (closure).
func Closure(a int) func() int {
	return func() int {
		return a
	}
}

// 4 - create a function that has two params of int an return the addition of it.
func Add(a, b int) int {
	return a + b
}
