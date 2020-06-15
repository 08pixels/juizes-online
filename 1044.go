package main

import "fmt"

var numberA, numberB int

func main() {
	fmt.Scan(&numberA, &numberB)
	if numberA%numberB == 0 || numberB%numberA == 0 {
		fmt.Println("Sao Multiplos")
	} else {
		fmt.Println("Nao sao Multiplos")
	}
}
