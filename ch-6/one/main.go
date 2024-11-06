// Create a struct named Person with three fields: FirstName and LastName of type string and
// Age of type of type int. Write a function called MakePerson that takes in firstName, lastName,
// and returns a Person. Write a secind function MakePersonPointer that takes in firstName, lastName,
// and age and returns a *Person. Call both from main. Compile your program with go build -gcflags="-m".
// This both compiles your code and prints out which values escapes to the heap. Are you surprised about
// what escapes?

package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

// func returnPointer[T any](t T) *T {
// 	return &t
// }

func MakePerson(firstName, lastName string, age int) Person {
	return Person{firstName, lastName, age}
}

func MakePersonPointer(firstName, lastName string, age int) *Person {
	return &Person{firstName, lastName, age}
}

func main() {
	person := MakePerson("Talut", "Salako", 25)
	personPointer := MakePersonPointer("Talut", "Salako", 25)

	fmt.Println(person)
	fmt.Println(personPointer)
}
