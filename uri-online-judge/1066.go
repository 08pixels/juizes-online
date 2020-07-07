package main

import "fmt"

var number int

var even int
var odd int
var positive int
var negative int

func main() {
	i := 0
	for i < 5 {
		fmt.Scan(&number)

		if number%2 == 0 {
			even++
		} else {
			odd++
		}

		if number > 0 {
			positive++
		} else if number < 0 {
			negative++
		}

		i++
	}

	fmt.Printf("%d valor(es) par(es)\n", even)
	fmt.Printf("%d valor(es) impar(es)\n", odd)
	fmt.Printf("%d valor(es) positivo(s)\n", positive)
	fmt.Printf("%d valor(es) negativo(s)\n", negative)
}
