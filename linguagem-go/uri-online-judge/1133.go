package main

import "fmt"

var inf, sup int

func main() {
	fmt.Scan(&inf, &sup)

	if inf > sup {
		inf, sup = sup, inf
	}

	for i := inf + 1; i < sup; i++ {
		if i%5 == 2 || i%5 == 3 {
			fmt.Println(i)
		}
	}

}
