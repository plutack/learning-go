// Write a for loop rhar puts 100 random numbers between 0 and 100 into int slice.

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	slice := make([]int, 0, 100)
	for range 100 {
		randInt := rand.Intn(100)
		slice = append(slice, randInt)
		fmt.Println(slice)
	}
	// Loop over the slice you created in ex 1. For each value in the
	// slice, apply the following rules:
	// a. If the value is divisible by 2, print "Two!"
	// b. If the value is divisible by 3, print "Three!"
	// c. If the value is divisible by 2 and 3, Don't print anything else..
	// d. Otherwise, print "Never mind".

	for _, v := range slice {
		switch {
		case v%2 == 0 && v%3 == 0:
		case v%2 == 0:
			fmt.Println("Two!")
		case v%3 == 0:
			fmt.Println("Three!")

		default:
			fmt.Println("Never mind")
		}
	}
}
