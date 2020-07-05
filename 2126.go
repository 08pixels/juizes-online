package main

import (
	"fmt"
)

func main() {
	var cases int
	var sub, str string

	for true {
		_, err := fmt.Scan(&sub, &str)

		if err != nil {
			break
		}

		amount := 0
		position := 0

		for i := 0; i < len(str); i++ {

			counter := 0
			for j := 0; j < len(sub); j++ {
				if i+j < len(str) && str[i+j] == sub[j] {
					counter++
				}
			}
			if counter == len(sub) {
				position = i + 1
				amount++
				i = i + len(sub) - 1
			}
		}
		cases++
		fmt.Printf("Caso #%d:\n", cases)

		if amount == 0 {
			fmt.Printf("Nao existe subsequencia\n\n")
		} else {
			fmt.Printf("Qtd.Subsequencias: %d\n", amount)
			fmt.Printf("Pos: %d\n\n", position)
		}
	}
}
