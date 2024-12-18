package main

import "fmt"

func main() {
	s := make([]func(), 4)

	for i := 0; i < 4; i++ {
		s[i] = func() {
			fmt.Printf("%d @ %p\n ", i, &i)
		}
	}

	for i := 0; i < 4; i++ {
		s[i]()
	}
}
