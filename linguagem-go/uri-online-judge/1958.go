package main

import (
	"fmt"
	"strconv"
)

var number string

func main() {
	fmt.Scan(&number)
	value, _ := strconv.ParseFloat(number, 64)

	if number[0] == '-' {
		fmt.Printf("%.4E\n", value)
	} else {
		fmt.Printf("+%.4E\n", value)
	}
}
