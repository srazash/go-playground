package main

import "fmt"

func main() {
	for count := 1; count <= 100; count++ {
		switch {
		case count%15 == 0:
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
