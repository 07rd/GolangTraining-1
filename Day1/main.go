package main

import "fmt"

func main() {

	fmt.Println("Assignment 1")
	fmt.Println("Hello World!")
	fmt.Println("===============================================\n")
	fmt.Println("Assignment 2")
	pyramid(5, true)
	fmt.Println("Assignment 3")
	pyramid(5, false)
	fmt.Println("Assignment 4")
	factorialNum(5)
	fmt.Println("===============================================\n")
	fmt.Println("Assignment 5")
	isPrime(11)
}

func pyramid(rowNumber int, astric bool) {

	if (astric) {
		for row := 0; row < rowNumber; row++ {
			for col := 0; col <= row; col++ {
				fmt.Print("* ")
			}
			fmt.Println()
		}
	} else {
		var num int = 1
		for row := 0; row < rowNumber; row++ {
			for col := 0; col <= row; col++ {
				fmt.Print(fmt.Sprint(num, " "))
				num++
			}
			fmt.Println()
		}	
	}
	fmt.Println("===============================================\n")
}

func factorialNum(num int) {
	var value = 1
	for i := 1; i <= num; i++ {
			value *= i
		}
	fmt.Println(fmt.Sprint("Factorial of number ", num, " is ", value))
}

func isPrime(num int) {
	if num < 2 {
		fmt.Println(fmt.Sprint("Not a prime number ", num))
		return
	}
	for i := 2; i <= num/2; i++ {
		if num%i == 0 {
			fmt.Println(fmt.Sprint("Not a prime number ", num))
			return
		}
	}
	fmt.Println(fmt.Sprint(num, " is a prime number "))	
}
