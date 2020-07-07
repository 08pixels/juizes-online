package main

import "fmt"

var cases int
var number int
var note float64

func main() {
	fmt.Scan(&cases)

	answerNumber := -1
	answerNote := float64(0.0)

	for cases != 0 {
		fmt.Scan(&number, &note)

		if note >= 8.0 && note > answerNote {
			answerNumber = number
			answerNote = note
		}
		cases--
	}

	if answerNumber == -1 {
		fmt.Println("Minimum note not reached")
	} else {
		fmt.Println(answerNumber)
	}
}
