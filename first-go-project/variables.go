package main

import "reflect"

func Variables(){
	var firstNumber int
	var firstFloat float32
	firstNumber, firstFloat = 2, 2.5
	// second type of variable declaration
	secondNumber := 58.5
	println(reflect.TypeOf(secondNumber).String())
	// print address
	println(&secondNumber)
	println(firstNumber, firstFloat)
	
}