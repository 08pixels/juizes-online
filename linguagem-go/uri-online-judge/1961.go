package main

import "fmt"

var jump, amount int
var last, number int

func abs(number int) int {

	if number < 0 {
		return -number
	}

	return number
}

func main() {
	fmt.Scan(&jump, &amount)
	answer := "YOU WIN"

	for i := 0; i < amount; i++ {
		fmt.Scan(&number)

		if i != 0 {
			if abs(number-last) > jump {
				answer = "GAME OVER"
			}
		}

		last = number
	}
	fmt.Println(answer)
}
