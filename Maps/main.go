package main

import "fmt"

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

	delete(studentsGrade, "Dipen") // deleting the values
	// fmt.Println("New marks of Dipen is:", studentsGrade["Dipen"])   //checking if value is deleted

	fmt.Println("Marks of Pankaj is", studentsGrade["Pankaj"])

	// check is a student exists or not OR if his marks are added or not
	grades, exists := studentsGrade["Pankaj"]
	fmt.Println("Grades of Pankaj is:", grades)
	fmt.Println("Pankaj exists:", exists)
}
