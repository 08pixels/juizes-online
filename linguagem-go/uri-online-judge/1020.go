package main

import "fmt"

var ageInDays int

func main() {
	fmt.Scanf("%d", &ageInDays)
	years := ageInDays / 365
	ageInDays %= 365
	months := ageInDays / 30
	fmt.Printf("%d ano(s)\n%d mes(es)\n%d dia(s)\n", years, months, ageInDays%30)
}
