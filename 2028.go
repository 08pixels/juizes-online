package main

import "fmt"

var number int
var cases int

func main() {

	amount := [201]int{}
	amount[0] = 1
	for i := 1; i <= 200; i++ {
		amount[i] = amount[i-1] + i
	}

	for true {
		_, err := fmt.Scan(&number)

		if err != nil {
			break
		}

		cases++
		if amount[number] == 1 {
			fmt.Printf("Caso %d: %d numero\n", cases, amount[number])
		} else {
			fmt.Printf("Caso %d: %d numeros\n", cases, amount[number])
		}

		for i := 0; i <= number; i++ {
			if i == 0 {
				fmt.Printf("%d", i)
			}
			for j := 1; j <= i; j++ {
				fmt.Printf(" %d", i)
			}
		}
		fmt.Printf("\n\n")
	}
}
