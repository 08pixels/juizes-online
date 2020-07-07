package main

import "fmt"

var cases, number int
var lowest, position int

func main() {

	fmt.Scan(&cases)
	for i := 0; i < cases; i++ {
		fmt.Scan(&number)

		if lowest > number || i == 0 {
			lowest = number
			position = i
		}
	}
	fmt.Printf("Menor valor: %d\n", lowest)
	fmt.Printf("Posicao: %d\n", position)
}
