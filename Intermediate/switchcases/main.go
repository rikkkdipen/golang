package main

import (
	"fmt"
)

func main() {
	day := 2
	month := "January"
	temperature := 25

	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	default:
		fmt.Println("Unknown Day")
	}

	switch month {
	case "January", "february", "March":
		fmt.Println("Winter")
	case "April", "May", "June":
		fmt.Println("Summer")
	case "July", "August", "September":
		fmt.Println("Rainy")
	}

	switch {
	case temperature < 0:
		fmt.Println("Freezy")
	}

}
