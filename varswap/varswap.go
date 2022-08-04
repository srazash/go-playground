package main

import (
	"fmt"
)

func main() {
	var a = 87
	var b = 194

	fmt.Println("Variables before swap:")
	fmt.Println(a, b)

	// swap variabled without using a third variable
	a = a + b
	b = a - b
	a = a - b

	fmt.Println("Variables after swap:")
	fmt.Println(a, b)
}
