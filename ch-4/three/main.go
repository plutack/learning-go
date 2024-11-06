// Start a new program. In main, declare an int variable called total. Write a for loop that
// uses a variable named i to iterate from 0 (inclusive) to 10 (inclusive). The body of the
// loop should be as follows:
// total := total + 1
// fmt.Println(total)
// After the loop, print out the value of total, What is printed out? What is the likely bug in
// this code ?

package main

import "fmt"

func main() {
	var total int
	for i := 0; i <= 10; i++ {
		// total := total + 1 // total is reinitialized each time also shadowed?
		total = total + 1
		fmt.Println(total)
	}
}
