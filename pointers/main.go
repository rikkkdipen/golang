package main

import "fmt"

func modifyingValueByReference(num *int) {
	*num = *num + 20
}

func main() {

	// var num int = 2
	// var ptr *int
	// ptr = &num

	num := 2
	ptr := &num

	// fmt.Println("Num has value", num)
	fmt.Println("Pointer contains", ptr)
	fmt.Println("Data contains through pointer is", *ptr)
	// *int means the data to which it is pointed to contains the data in the form of int

	var pointer *int
	if pointer == nil {
		fmt.Println("Pointer is nil")
	} else {
		fmt.Println("Pointer contains a value")
	}

	value := 5
	modifyingValueByReference(&value)
	fmt.Println("value Contains: ", value)
}
