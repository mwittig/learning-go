package main

import (
	"fmt"
)

func getSumAndProductOf(x, y int) (sum, prod int) {
	// it is possible to return multiple values separated by comma
	return x + y, x * y
}

func main () {
	// := is a type of short declaration to infer the type from the given expression
	x := 4
	y := 3
	sum, prod := getSumAndProductOf(x, y)
	fmt.Println("sum:", sum, "product:", prod)
}
