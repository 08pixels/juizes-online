package main

import "fmt"

var cases int
var item, amount int
var total float64

func main() {
	fmt.Scan(&cases)

	for cases != 0 {
		fmt.Scan(&item, &amount)

		if item == 1001 {
			total += (1.50 * float64(amount))
		} else if item == 1002 {
			total += (2.50 * float64(amount))
		} else if item == 1003 {
			total += (3.50 * float64(amount))
		} else if item == 1004 {
			total += (4.50 * float64(amount))
		} else {
			total += (5.50 * float64(amount))
		}
		cases--
	}

	fmt.Printf("%.2f\n", total)
}
