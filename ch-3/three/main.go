// Write  a program that defines a struct called Employee with three fields: firstname,
// lastname and id. The first two fields of type string, and the last field is of type
// int. Initialize the first  one using the struct literal style without names, the second
// using the struct literal style with names, and the third with a var declaration. Use dot
// notation to populate the fileds in the thied struct. Print out the three structs

package main

import "fmt"

type Employee struct {
	firstname string
	lastname  string
	id        int
}

func main() {
	firstVar := Employee{
		firstname: "Jane",
		id:        1,
		lastname:  "Doe",
	}
	secondVar := Employee{
		"Talut",
		"Salako",
		2,
	}
	var thirdVar Employee

	thirdVar.firstname = "John"
	thirdVar.lastname = "Doe"
	thirdVar.id = 3

	fmt.Printf("the variables are %v, %v, %v", firstVar, secondVar, thirdVar)
}
