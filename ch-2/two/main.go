// write a program that declares a constant called valur that can be assigned to both
// an integer and a floating-point variable. Assign it to an integer called i and a
// floating-point called f. Print out i and f.

package main

import "fmt"

func main() {
	const value = 5

	i := int(value)
	f := float64(value)

	fmt.Printf("i is of type: %[1]T and value: %[1]v\n", i)
	fmt.Printf("i is of type: %[1]T and value: %[1]v\n", f)
}
