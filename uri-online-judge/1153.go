package main

import "fmt"

func factorial(number int) int {
	result := number

	for i := 2; i < number; i++ {
		result *= i
	}

	return result
}

func main() {
	var number int
	fmt.Scan(&number)
	fmt.Println(factorial(number))
}
