package main

import (
	"fmt"
	"main/myutil"
)

func main() {

	fmt.Println("This code is written by Dipen")
	myutil.Printmessage("Hello Hii") // calling another function or package from different folder

	var name string = "Dipen"
	fmt.Println(name)

	var version = "V1.latest"
	fmt.Println(version)

	var money int = 500000
	var currency = 5000
	fmt.Println(money)
	fmt.Println("Currency is ", currency)

	var dimension float64 = 87.13
	fmt.Println(dimension)

	var decided bool = false
	fmt.Println(decided)

	var person = 23
	fmt.Println(person)

	const pi = 3.14 // const values cannot be updated but var values can
	fmt.Println(pi)

	person1 := "Dipen Rikame"
	fmt.Println(person1)

	var Public = "Data is important"
	var private = "Data is private"

	fmt.Println(Public)  // for public variable keep first letter in uppercase()
	fmt.Println(private) // for private variable keep first letter in lowercase()

	// func Publicfunction(){
	// 	fmt.Println("This is a public function")  // same public and private rule is for functions
	// }

	// func privatefucntion(){
	// 	fmt.Println("This is a private function")  // refer the above comment
	// }

}
