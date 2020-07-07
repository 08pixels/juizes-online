package main

import "fmt"

var cases, number int

func main() {
	fmt.Scan(&cases)

	for cases != 0 {
		fmt.Scan(&number)
		fmt.Println(number & 1)
		cases--
	}
}
