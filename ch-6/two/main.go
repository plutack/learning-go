// Write two functions. The UpdateSlice function that takes in a [] string and a string.
// It sets the last position in the passed-in slice to the passed-in string. At the end of
// UpdateSlice, print the slice after making the change. The GrowSlice function also takes in a
// [] string and a string. It appends the string onto the slice. At the end of GrowSlice, print
// the slice after making the change. Call these functions from main. Print out the slice before
// each function is called  and after each function is called. Do you understand why some changes
// are visible in main and why some changes are not?

package main

import "fmt"

func UpdateSlice(arr []string, updateValue string) {
	lastIndex := len(arr) - 1
	arr[lastIndex] = updateValue
	fmt.Println(arr)
}

func GrowSlice(arr []string, appendValue string) {
	arr = append(arr, appendValue)
	fmt.Println(arr)
}

func main() {
	arr := []string{"Talut", "Korede", "Salako"}
	fmt.Println(arr)
	updateValue := "plutack"
	UpdateSlice(arr, updateValue)
	fmt.Println(arr)
	GrowSlice(arr, updateValue)
	fmt.Println(arr)
}
