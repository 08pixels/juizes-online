package main

import (
	"fmt"
	"strconv"
)

func main() {
	var number string

	for true {
		fmt.Scan(&number)

		if number == "-1" {
			break
		}

		if len(number) > 1 && number[1] == 'x' {
			value, _ := strconv.ParseInt(number[2:], 16, 32)
			fmt.Printf("%d\n", value)
		} else {
			value, _ := strconv.ParseInt(number, 10, 32)
			fmt.Printf("0x%X\n", value)
		}
	}
}
