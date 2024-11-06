// Write a function called prefixer that has  an input parameter of type string and returns a function that has a parameter
// of type string and returns a string. The returned function should prefix its input with the string passed into Prefixer.
// Use the following main function to test prefixer.

package main

import "fmt"

type retFunc func(string) string

var r retFunc = func(s string) string {
	return s
}

func prefixer(prefix string) retFunc {
	return func(s string) string {
		return fmt.Sprintf("%s %s", prefix, s)
	}
}

func main() {
	pref := prefixer("hello")
	fmt.Println(pref("Talut"))
}
