package main

import "fmt"

func main() {

	var count int = 1

	for count <= 100 {
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
		count++
	}

}
