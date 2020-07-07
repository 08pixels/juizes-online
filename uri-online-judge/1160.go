package main

import "fmt"

var cases int
var popA, popB int
var rateA, rateB float64

func main() {
	fmt.Scan(&cases)

	for cases != 0 {
		fmt.Scan(&popA, &popB, &rateA, &rateB)
		rateA /= 100
		rateB /= 100
		i := 0

		for i < 100 {
			popA += int(float64(popA) * rateA)
			popB += int(float64(popB) * rateB)

			if popA > popB {
				fmt.Printf("%d anos.\n", i+1)
				break
			}
			i++
		}

		if i == 100 {
			fmt.Println("Mais de 1 seculo.")
		}

		cases--
	}
}
