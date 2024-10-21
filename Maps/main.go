package main

// maps are alwasys in unordered way
import (
	"fmt"
)

func main() {

	fmt.Println("This is Maps in Golang")

	// Maps are used to associate values with keys

	studentsGrade := make(map[string]int)

	studentsGrade["Dipen"] = 50
	studentsGrade["Deven"] = 25
	studentsGrade["Pankaj"] = 48
	studentsGrade["Darpan"] = 49
	studentsGrade["Prashant"] = 45
	studentsGrade["Sangeet"] = 40

	fmt.Println("Marks of Dipen:", studentsGrade["Dipen"]) // recalling the values
	studentsGrade["Dipen"] = 100                           // updating the values
	fmt.Println("New marks of Dipen is:", studentsGrade["Dipen"])

	// delete(studentsGrade, "Dipen") // deleting the values
	// fmt.Println("New marks of Dipen is:", studentsGrade["Dipen"])   //checking if value is deleted

	fmt.Println("Marks of Pankaj is", studentsGrade["Pankaj"])

	// check is a student exists or not OR if his marks are added or not
	grades, exists := studentsGrade["Pankaj"]
	fmt.Println("Grades of Pankaj is:", grades)
	fmt.Println("Pankaj exists:", exists)

	for index, value := range studentsGrade { //using for loop for returning values

		fmt.Printf("Key is %s and marks is %d\n", index, value) // %s for string and %d for integer value == marks
	}

	person := map[string]int{

		"Alice":   90,
		"Bob":     45,
		"Charlie": 95,
	}

	for index, value := range person {

		fmt.Printf("------Key is: %s and value is: %d\n", index, value)
	}

}
