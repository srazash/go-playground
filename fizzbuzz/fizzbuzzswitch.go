package main

import (
	"fmt"
)

func main() {
	for count := 1; count <= 1000; count++ {
		switch {
		//this case could also be count%15 == 0, but spliting this is more idiomatic
		case count%5 == 0 && count%3 == 0:
			fmt.Println(count, " FizzBuzz!")
		case count%5 == 0:
			fmt.Println(count, " Buzz!")
		case count%3 == 0:
			fmt.Println(count, " Fizz!")
		default:
			fmt.Println(count)
		}
	}
}
