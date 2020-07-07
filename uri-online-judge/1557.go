package main

import (
	"fmt"
	"math"
	"strconv"
)

var number int

func main() {
	for true {
		fmt.Scan(&number)

		if number == 0 {
			break
		}

		largestNumber := 1 << (2 * (uint32(number) - 1))
		amountOfDigits := int(math.Log10(float64(largestNumber))) + 1

		if number != 1 {
			amountOfDigits++
		}

		formated := "%" + strconv.Itoa(amountOfDigits) + "d"
		firstCol := "%" + strconv.Itoa(amountOfDigits-1) + "d"

		for i := 0; i < number; i++ {
			for j := 0; j < number; j++ {
				if j == number-1 {
					fmt.Printf(formated+"\n", 1<<(uint32(i+j)))
				} else if j == 0 {
					fmt.Printf(firstCol, 1<<(uint32(i+j)))
				} else {
					fmt.Printf(formated, 1<<(uint32(i+j)))
				}
			}
		}

		fmt.Printf("\n")
	}
}
