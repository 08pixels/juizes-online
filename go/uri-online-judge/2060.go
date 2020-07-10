package main

import "fmt"

func main() {
	var cases, number int
	var amount [6]int

	fmt.Scan(&cases)
	for cases != 0 {
		fmt.Scan(&number)

		for i := 2; i <= 5; i++ {
			if number%i == 0 {
				amount[i]++
			}
		}
		cases--
	}

	for i := 2; i <= 5; i++ {
		fmt.Printf("%d Multiplo(s) de %d\n", amount[i], i)
	}
}
