package main

import (
	"fmt"
	"time"
	// "golang.org/x/text/message"
)

// import (
// 	"fmt"
// 	"time"
// )

// import "github.com/plutack/basic-go/types"
// var mainNumber byte = 20

// func addOne(pointer *byte) {
// 	*pointer++
// 	fmt.Println(*pointer)
// 	// return number
// }

// go channels
func loop(word string, response *chan string, message string) {
	for i := 0; i <= 5; i++ {
		fmt.Println(i)
		time.Sleep(1 * time.Second)
	}
	*response <- message
}

func main() {
	response := make(chan string, 2)
	go loop("foo", &response, "first loop done")
	time.Sleep(5 * time.Second)
	go loop("bar", &response, "second loop done")
	fmt.Println(<-response)
	fmt.Println(<-response)
}

// func main() {
// 	// types.Run()
// 	types.Main()
// 	// var firstNumber, SecondNumber float32
// 	// var operation string
// 	// fmt.Println("Simple Calculator")
// 	// fmt.Println("================")
// 	// time.Sleep(1000)
// 	// fmt.Println("Enter operation")
// 	// fmt.Scanf("%s", &operation )
// 	// time.Sleep(3000)
// 	// fmt.Println("Enter first number")
// 	// fmt.Scanf("%f", &firstNumber )
// 	// time.Sleep(3000)
// 	// fmt.Println("Enter second number")
// 	// fmt.Scanf("%f", &SecondNumber )
// 	// time.Sleep(3000)
// 	// answer, error := calculator(firstNumber, SecondNumber, operation)
// 	// if error != nil {
// 	// 	fmt.Println(error)
// 	// } else {
// 	// 	fmt.Printf("The answer is %f", answer)
// 	// }
// 	// fmt.Printf(" answer is is %v", mainNumber)
// 	// println("Hello world")

// }

// func calculator (firstNumber float32, secondNumber float32, operation string) (float32, error){
// 		switch operation{
// 			case "add" :
// 				return firstNumber + secondNumber, nil
// 			case "subtract" :
// 				return firstNumber + secondNumber, nil
// 			case "multiply" :
// 				return firstNumber + secondNumber, nil
// 			case "divide" :
// 				if firstNumber == 0 {
// 					return 0.0, fmt.Errorf("first number cannot be zero")
// 				}
// 				return firstNumber + secondNumber, nil
// 			default:
// 				return 0.00, fmt.Errorf("invalid Operation")
// 		}

// 	}
