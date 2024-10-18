package main

import "fmt"

func main() {
	fmt.Println("This code is for arrays in Golang")

	// // 0,1,2,3,4
	// var name [5]string //made an array
	// name[0] = "Dipen"  // entered value in array
	// name[2] = "Deven"

	// fmt.Println("Names of Person is:", name)

	// var numbers = [5]int{1, 2, 3, 4, 5} // made an array and entered value at same time
	// fmt.Println("Numbers are:", numbers)

	// fmt.Println("length of numbers array is", len(numbers))
	// fmt.Println("value of name at 2nd index is", name[2])

	var price [5]string
	price[0] = "Dipen"
	price[2] = "Aakash"
	fmt.Println("The price is", price)
	fmt.Printf("Price array is %q\n", price)
}
