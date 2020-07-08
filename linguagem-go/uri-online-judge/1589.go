package main

import "fmt"

var cases int
var radiusA, radiusB int

func main() {
	fmt.Scan(&cases)

	for cases != 0 {
		fmt.Scan(&radiusA, &radiusB)
		fmt.Println(radiusA + radiusB)
		cases--
	}
}
