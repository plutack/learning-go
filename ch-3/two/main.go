// Write a program that defines a string variable called message with the value
// "Hi 👩 and 👨" and print the fourth rune as a character, not a number
package main

import "fmt"

func main() {
	message := "Hi 👩 and 👨"
	msgAsRunes := []rune(message)

	fmt.Printf("Third char as string is: %s", string(msgAsRunes[3]))
}
