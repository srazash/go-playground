/*
	Based on some VERY old code I wrote in Java which prints a letter N	made
	out of multiple Ns based on the input from the command line args.
*/

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	sctmp := ""
	scale := 5 // default scale is 3

	// scale is only changed if any user args exist
	if len(os.Args) > 1 {
		sctmp = os.Args[1]
		scale, _ = strconv.Atoi(sctmp)
	}

	for line := 0; line < scale; line++ {
		fmt.Print("N")
		for column := 1; column < scale-1; column++ {
			switch {
			case column == line:
				fmt.Print("N")
			default:
				fmt.Print(" ")
			}
		}
		fmt.Println("N")
	}
}
