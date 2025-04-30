package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	fmt.Println("Enter a name:")
	// var name string

	// fmt.Scan(&name) //one way to take input
	// fmt.Println("Hello, Mr.", name)

	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	fmt.Println("Hello, Mr.", name)

}
