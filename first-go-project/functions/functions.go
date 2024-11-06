package functions

import "fmt"

// how variables are affected



// func addOne(number byte) byte {
// 	number ++
// 	fmt.Println(number)
// 	// return number 
// }


func AddOne(pointer *byte) {
	*pointer ++
	fmt.Println(&pointer)
	// return number
}

