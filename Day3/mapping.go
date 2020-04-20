package main

import "fmt"

func studentInfo() {
	studentMap := make(map[int]string)
	studentMap[1] = "Ron"
	studentMap[2] = "Stella"
	studentMap[3] = "Amelia"
	studentMap[4] = "Augusta"
	studentMap[5] = "Jack"

	for key, val := range studentMap{
		fmt.Printf("Roll no. %d : %s\n", key, val)
	}

	updateName(studentMap,"Kevin", 2)
	fmt.Println("\nAfter name updated: ", studentMap)

	removeName(studentMap, 4)
	fmt.Println("\nAfter removal: ",studentMap)
}

func updateName(student_map map[int]string, name string, rollNumber int){
	student_map[rollNumber] = name
}

func removeName(student_map map[int]string, rollNumber int) {
	delete(student_map, rollNumber)
}