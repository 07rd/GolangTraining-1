package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	fmt.Println("Assignment 1:\n")
	hitURL("https://tour.golang.org")
	fmt.Println("Assignment 2:\n")
	studentInfo()
	fmt.Println("Assignment 3:\n")
	sliceOperations()
}

func hitURL(url string) int {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error in response. :", err)
		return resp.StatusCode
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error in Body ", err)
			return 0
		}
		bodyString := string(bodyBytes)
		
		fmt.Println("Body: ", bodyString)
	}
	return resp.StatusCode
}