package main

import (
	"fmt"
)

var inf, sup, sumOdd int

func main() {
	fmt.Scan(&inf, &sup)

	if inf > sup {
		inf, sup = sup, inf
	}

	for i := inf + 1; i < sup; i++ {
		if i%2 != 0 {
			sumOdd += i
		}
	}
	fmt.Println(sumOdd)
}
