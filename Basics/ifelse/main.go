package main

import "fmt"

func main() {
	x := 10
	if x > 5 {
		fmt.Println("x is greater than 5")
	} else {
		fmt.Println("x is less than 5")
	}

	z := 10
	if z > 10 {
		fmt.Println("Z is greater than 10")
	} else if z > 5 {
		fmt.Println("Z is greater than 5 but less than 10")
	} else {
		fmt.Println("Z is smalller tha 5")
	}

	y := 10
	if x > 5 && (y > 5 || z == 5) {
		fmt.Println("Hii how are you")
	} else {
		fmt.Println("learn programming with Dipen")
	}
}
