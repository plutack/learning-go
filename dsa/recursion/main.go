package main

import "fmt"

func main() {
	fmt.Println(reverseString("Hello"))
}

func reverseString(s string) string {
	if len(s) == 1 {
		return string(s[0])
	}

	return string(s[len(s)-1]) + reverseString(s[:len(s)-1])
}
