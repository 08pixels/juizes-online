package main

import "fmt"

var cases, force int
var name string

func main() {
	fmt.Scan(&cases)
	for cases != 0 {
		fmt.Scan(&name, &force)
		if name == "Thor" {
			fmt.Println("Y")
		} else {
			fmt.Println("N")
		}
		cases--
	}
}
