package main

import (
	"fmt"
)

func main() {
	var a, b float32

	fmt.Println("Input a: ")
	fmt.Scan(&a)

	fmt.Println("Input b: ")
	fmt.Scan(&b)

	fmt.Println("Result =: ", (a+b)/2)
}
