package main

import "fmt"

var cases int
var number string

func main() {
	fmt.Scan(&cases)

	for cases != 0 {
		fmt.Scan(&number)

		answer := 0
		for i := range number {

			switch number[i] {
			case '1':
				answer += 2
			case '2', '3', '5':
				answer += 5
			case '4':
				answer += 4
			case '7':
				answer += 3
			case '8':
				answer += 7
			default:
				answer += 6
			}
		}
		fmt.Printf("%d leds\n", answer)
		cases--
	}
}
