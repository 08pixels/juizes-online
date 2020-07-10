package main

import "fmt"

var valueA, valueB float64

func main() {
	fmt.Scan(&valueA, &valueB)

	dif := valueB - valueA

	fmt.Printf("%.2f%%\n", (dif*100)/valueA)
}
