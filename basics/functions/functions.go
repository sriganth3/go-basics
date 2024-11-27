package main

import "fmt"

// var statement can be at package or function level
var c, python, java bool

// var declarations can include initializers and type can be omitted if
// the initializer is present and variable will use the type of initializer

var j, k int = 1, 2

func main() {
	fmt.Println(add(5, 6))
	fmt.Println(sub(5, 6))
	a, b := swap("world", "hello")
	fmt.Println(a, b)
	fmt.Println(split(8))

	var i int
	fmt.Println(i, c, java, python)

	var c, python, java = true, false, "yes!"
	fmt.Println(j, k, c, python, java)

	// short variable assignment, outside a function, we need to use var. but inside the function := can be used
	l := 5
	fmt.Println(l)
}

func add(x int, y int) int {
	return x + y
}

// if the data type of the arguments is same continously, we can ignore giving the data types
func sub(x, y int) int {
	return x - y
}

// returning multiple return params
func swap(x, y string) (string, string) {
	return y, x
}

// return variable can be defined on the function level. known as naked return
func split(sum int) (x, y int) {
	x = sum / 2
	y = sum - x
	return
}
