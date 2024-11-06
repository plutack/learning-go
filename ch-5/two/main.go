// Write a function called fileLen that has an input parameter of type string and returns an int and an error.
// The function takes in a filename and return sthe  number of bytes in the file. If there is an error reading
// the file, return the error. Use defer to make sure the file is closed properly.
package main

import (
	"fmt"
	"os"
)

func fileLen(input string) (*int, error) {
	data, err := os.Open(input)
	if err != nil {
		return nil, err
	}
	defer data.Close()
	b := make([]byte, 12000)
	// could have used stat() file method
	length, err := data.Read(b)
	if err != nil {
		return nil, err
	}
	return &length, nil
}

func main() {
	length, err := fileLen("/home/plutack/.zshrc")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(*length)
}
