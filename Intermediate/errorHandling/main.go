package main

import "fmt"

func divide(a, b float64) (float64, string) {
	if b == 0 {
		return 0, "Denominator must not be zero"
	}
	return a / b, "nil"
}

func main() {

	fmt.Println("This is error handling in go")
	ans, _ := divide(10, 5)
	fmt.Println("Division of 2 number is", ans)

	// fmt.Println("This is error handling in go")
	// ans, err := divide(10, 4)
	// if err != nil {
	// 	fmt.Println("Error handling")
	// }
	// fmt.Println("Division of 2 number is", ans)
}
