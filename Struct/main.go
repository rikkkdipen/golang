package main

import "fmt"

type Person struct {
	Firstname string
	Lastname  string
	Age       int
}

type Contact struct {
	Email string
	Phone string
}

type Address struct {
	Housenumber string
	Area        string
	State       string
}

type Employee struct {
	Person_details Person
	Person_Contact Contact
}

func main() {

	// 1st method
	// var dipen Person
	// fmt.Println("Dipen person:", dipen) // by default stored is 0

	// dipen.Firstname = "Dipen"
	// dipen.Lastname = "Rikame"
	// dipen.Age = 24
	// fmt.Println("Dipen person:", dipen)

	// 2nd method
	person1 := Person{
		Firstname: "Dipen",
		Lastname:  "Rikame",
		Age:       24,
	}

	fmt.Println("Person1 is:", person1)

	// 3rd method
	var person2 = new(Person)
	person2.Firstname = "Deven"
	person2.Lastname = "Rikame"
	person2.Age = 23

	fmt.Println("Person2:", person2) //& is used for pointer

	var employee1 Employee
	employee1.Person_details = Person{
		Firstname: "Dipen",
		Lastname:  "Rikame",
		Age:       24,
	}
	employee1.Person_Contact.Email = "contact@gmail.com"
	employee1.Person_Contact.Phone = "7400173874"

	fmt.Println("Employee1", employee1)

}
