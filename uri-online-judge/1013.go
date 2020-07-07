package main

import (
	"fmt"
)

var (
	valueA int
	valueB int
	valueC int
)

func max(valueA int, valueB int) int {
	if valueA > valueB {
		return valueA
	}
	return valueB
}

func main() {
	fmt.Scanf("%d %d %d", &valueA, &valueB, &valueC)
	fmt.Printf("%d eh o maior\n", max(max(valueA, valueB), valueC))
}
