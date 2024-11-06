// Write a program that declares an integer  variable called i with a value 20.
// Assign i to a floating-point variable named f. Print out  i and f

package main

import "fmt"

func main() {
	i := 20

	f := float64(i)

	fmt.Printf("i is of type: %[1]T and value: %[1]d\nf is of type: %[2]T and value: %[2]f", i, f)
}
