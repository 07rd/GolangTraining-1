package main

import "fmt"

func sliceOperations() {
	var text_slice =[]string{"This", "is", "a", "Go", "Language"}

	fmt.Println("Length of slice is : ", len(text_slice))

	fmt.Println("Before update : ", text_slice)
	text_slice[2] = "not"
	fmt.Println("After update : ", text_slice)

	newText := make([]string, len(text_slice))
	copy(newText, text_slice)
	newText[3] = "Java"
	fmt.Println("New slice after update: ", newText)
	fmt.Println("Old slice after update: ", text_slice)
}
