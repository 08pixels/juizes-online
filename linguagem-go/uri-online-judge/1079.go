package main

import "fmt"

var cases int
var a, b, c float64

func main() {
	fmt.Scan(&cases)

	for i := 0; i < cases; i++ {
		fmt.Scan(&a, &b, &c)
		fmt.Printf("%.1f\n", (2*a+3*b+5*c)/10)
	}
}
