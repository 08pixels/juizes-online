package main

import "fmt"

var cases, number int

func isPerfect(number int) bool {
	divSum := 1
	for i := 2; i*i <= number; i++ {
		if number%i == 0 {
			divSum += i

			if i*i != number {
				divSum += (number / i)
			}
		}
	}

	return divSum == number && number != 1
}

func main() {
	fmt.Scan(&cases)

	for cases != 0 {
		fmt.Scan(&number)

		if isPerfect(number) {
			fmt.Printf("%d eh perfeito\n", number)
		} else {
			fmt.Printf("%d nao eh perfeito\n", number)
		}
		cases--
	}
}
