package main

import "fmt"

func main() {

	age := 23
	name := "Dipen"
	height := 5.100

	// fmt.Println(name, age)
	// fmt.Println("Hello world")                // Println inserts a new line
	// fmt.Printf("age is %d\n", age)            // %d is format specifier for integer
	// fmt.Printf("height is %.2f\n", height)    // %f is format specifier for float
	// fmt.Printf("Type of age is %T\n", height) // %T is format specifier for Type
	// fmt.Printf("Name is %s\n", name)          // %s is format specifier for string

	fmt.Printf("Name is %s\n", name)
	fmt.Printf("age is %d\n", age)

	fmt.Printf("name is %s, Age is %d and height is %.2f\n", name, age, height)
}
