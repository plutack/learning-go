package main

import (
	"fmt"
	"time"
)

type Employee struct {
	Name  string
	Age   int
	Boss  *Employee
	Hired time.Time
}

func main() {
	e := Employee{
		Name:  "Talut",
		Age:   25,
		Boss:  nil,
		Hired: time.Now(),
	}
	b := Employee{"Bayo", 30, nil, time.Now()}
	e.Boss = &b
	fmt.Printf("%T %+[1]v\n", e)
	fmt.Printf("%T %+[1]v\n", b)
}
