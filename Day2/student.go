package main


/*
Programming assignment is as below:

1. Create struct for Student and create print method to print it. Include contact details struct in it.
2. Update Some information of that student. Use pointer.
Write Update function for it.

3. Test cases for checking contact details of student and age.

4. Save student info into file.
*/

import (
	"fmt"
	"io/ioutil"
	"os"
)

// Contacts
type Contacts struct {
	phone string
	email string
}

// Student
type Student struct {
	age  int
	name string
	Contacts
}

func (stud Student) print() {
	fmt.Printf("%+v\n", stud)
}

func (stud *Student) updateAge(newAge int) {
	stud.age = newAge
}

func (stud Student) saveToFile(filename string) {
	lines := fmt.Sprintf("%+v\n", stud)
	err := ioutil.WriteFile(filename, []byte(lines), 0666)
	if err != nil {
		os.Exit(1)
	}
}

func main() {
	s1 := Student{
		name: "Harry",
		age:  21,
		Contacts: Contacts{
			phone: "273971090", email: "harry.potter@sample.com",
		},
	}

	fmt.Println("+++++++++++++++++++ Before update +++++++++++++++++++")
	s1.print()
	s1.updateAge(18)
	fmt.Println("+++++++++++++++++++ After update +++++++++++++++++++")
	s1.print()
	fmt.Println("====================================================")
	fmt.Println("+++++++++++++++++++ Save to file +++++++++++++++++++")
	s1.saveToFile("./student_info.txt")
	fmt.Println("====================================================")
}
