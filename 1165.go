package main

import "fmt"

var cases, number int

func isPrime(number int) bool {

	for i := 2; i*i <= number; i++ {
		if number%i == 0 {
			return false
		}
	}

	return number != 1
}

func main() {
	fmt.Scan(&cases)

	for cases != 0 {
		fmt.Scan(&number)

		if isPrime(number) {
			fmt.Printf("%d eh primo\n", number)
		} else {
			fmt.Printf("%d nao eh primo\n", number)
		}
		cases--
	}
}
