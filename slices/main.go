package main

import "fmt"

func main() {

	fmt.Println("We are learning Slices in Golang")
	numbers := []int{1, 2, 3, 4, 5}
	numbers = append(numbers, 2, 10, 4, 6, 8, 12, 24, 45)
	fmt.Println("Number:", numbers)
	fmt.Printf("Numbers has data type: %T\n", numbers)
	fmt.Println("Length", len(numbers))

	// initializing a 0 size array
	name := []string{}
	fmt.Println("name is", name)

	// slice through make function

	numberss := make([]int, 3, 5)

	numberss = append(numberss, 4)
	numberss = append(numberss, 98)
	numberss = append(numberss, 6)

	fmt.Println("Slice", numberss)
	fmt.Println("capacity", cap(numberss))

	// var stock = []string{}
	// stock = append(stock, "Dipen")
	// fmt.Println(stock)

	//another way of writing with make method

	stock := make([]string, 0)
	stock = append(stock, "Dipen")
	stock = append(stock, "Deven")
	fmt.Println(stock)
}
