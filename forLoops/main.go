package main

import "fmt"

func main() {
	// In Golang for loop is the only way to use loops

	for i := 0; i < 6; i++ {
		fmt.Println("Numbers are:", i)
	}

	// use break statement

	counter := 0
	for {
		fmt.Println("Infinite loop")
		counter++
		if counter == 3 {
			break
		}
	}
}
