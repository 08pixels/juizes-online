package main

import (
	"fmt"
)

var cases int
var inf, sup int

func main() {
	fmt.Scan(&cases)

	for i := 0; i < cases; i++ {
		fmt.Scan(&inf, &sup)

		result := 0
		if inf > sup {
			inf, sup = sup, inf
		}

		for i := inf + 1; i < sup; i++ {
			if i%2 != 0 {
				result += i
			}
		}

		fmt.Println(result)
	}
}
