package main

import "fmt"

var cases int
var number int
var fibonnaci = [61]int64{}

func fillFibonnaci(value int) {
	fibonnaci[0] = 0
	fibonnaci[1] = 1

	for i := 2; i <= value; i++ {
		fibonnaci[i] = fibonnaci[i-1] + fibonnaci[i-2]
	}
}

func main() {
	fmt.Scan(&cases)

	fillFibonnaci(60)

	for cases != 0 {
		fmt.Scan(&number)
		fmt.Printf("Fib(%d) = %d\n", number, fibonnaci[number])
		cases--
	}
}
