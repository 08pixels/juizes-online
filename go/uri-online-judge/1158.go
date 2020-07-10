package main

import "fmt"

var cases int
var x, y int

func main() {

	fmt.Scan(&cases)

	for cases != 0 {
		fmt.Scan(&x, &y)

		if x&1 == 0 {
			x++
		}

		result := 0
		for i := 0; i < y; i++ {
			result += x
			x += 2
		}

		fmt.Println(result)
		cases--
	}

}
