// Write a small program with three variables, one named b of type byte, one named
// smallI of type int32, and one named bigI od type uint64. Assign each variale the maximum
// legal value for its type, an then add 1 to each  of ther variable. Print out their values

package main

import "fmt"

func main() {
	const addValue = 1

	b := ^byte(0)
	smallI := ^int32(0)
	bigI := ^uint64(0)

	fmt.Printf("old value: %v, %v, %v\n", b, smallI, bigI)
	b += byte(addValue)
	smallI += int32(addValue)
	bigI += uint64(addValue)
	fmt.Printf("new value: %v, %v, %v\n", b, smallI, bigI)
}
