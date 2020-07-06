package main

import "fmt"

func main() {
	var cases int
	var word string

	fmt.Scan(&cases)
	for cases != 0 {
		fmt.Scan(&word)
		fmt.Printf("%.2f\n", 0.01*float64(len(word)))
		cases--
	}
}
