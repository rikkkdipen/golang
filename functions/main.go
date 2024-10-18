package main

import "fmt"

func simpleFunction() {
	fmt.Println("This is a simple function")
}

func add(a, b int) int {
	return a + b

}

// another way of writing the same functin

// func add1(a, b int) (result int) {
// 	result = a + b
// 	return result

// }

func multiply(a, b int) int {

	return a * b
}

func main() {

	fmt.Println("We are learning functions in go")
	simpleFunction() // calling the function
	ans := add(3, 4)
	fmt.Println("Addition of two number is", ans)

	fmt.Println("Here we are going to multiply to numbers")
	result := multiply(3, 4)
	fmt.Println("The Multiplication of 2 numbers are", result)

}
