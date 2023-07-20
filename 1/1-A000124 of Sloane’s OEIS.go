package main

import (
	"fmt"
	"strconv"
)

func A000124(n int) []string {
	var sequence []string
	var current_number int = 1
	var increment int = 1

	for i := 1; i <= n; i++ {
		sequence = append(sequence, strconv.Itoa(current_number))
		current_number += increment
		increment++
	}

	return sequence
}

func main() {
	var input int

	for true {
		fmt.Print("Input: ")
		fmt.Scanln(&input)
		if input <= 0 {
			fmt.Println("Please insert positive number above zero")
		} else {
			fmt.Print("Output: ")
			for i, n := range A000124(input) {
				if (i != len(A000124(input))-1) {
					fmt.Print(n, "-")
				} else {
					fmt.Print(n)
				}
			}
			break
		}
	}
}
