package main

import "fmt"

var cases, year int

func main() {

	fmt.Scan(&cases)

	for cases != 0 {
		fmt.Scan(&year)
		if year >= 2015 {
			fmt.Printf("%d A.C.\n", year-2014)
		} else {
			fmt.Printf("%d D.C.\n", 2015-year)
		}
		cases--
	}
}
