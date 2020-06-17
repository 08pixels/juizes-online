package main

import "fmt"

var month = []string{
	"January",
	"February",
	"March",
	"April",
	"May",
	"June",
	"July",
	"August",
	"September",
	"October",
	"November",
	"December",
}

var ordinal int

func main() {
	fmt.Scan(&ordinal)
	fmt.Println(month[ordinal-1])
}
