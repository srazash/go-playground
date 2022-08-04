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
	scale := 5 // default scale is 5

	// scale is only changed if any user args exist
	// otherwise remains the default as initially declared
	if len(os.Args) > 1 {
		sctmp = os.Args[1]
		scale, _ = strconv.Atoi(sctmp)
	}

	// works line by line
	for line := 0; line < scale; line++ {
		fmt.Print("N") // first character of the line should always be an N

		// calculates where the N should be printed in the current line
		// to connect the first and last line of Ns diagonally
		for column := 1; column < scale-1; column++ {
			switch {
			case column == line:
				fmt.Print("N")
			default:
				fmt.Print(" ")
			}
		}
		fmt.Println("N") // last character of the line should always be an N, then a new line
	}
}
