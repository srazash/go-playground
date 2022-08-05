package main

import (
	"fmt"
)

func main() {
	for count := 1; count <= 1000; count++ {
		if count%5 == 0 && count%3 == 0 {
			fmt.Println(count, " FizzBuzz!")
		} else if count%5 == 0 {
			fmt.Println(count, " Buzz!")
		} else if count%3 == 0 {
			fmt.Println(count, " Fizz!")
		} else {
			fmt.Println(count)
		}
	}
}
