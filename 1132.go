package main

import "fmt"

var inf, sup, sum int

func main() {
	fmt.Scan(&inf, &sup)

	if inf > sup {
		inf, sup = sup, inf
	}

	for i := inf; i <= sup; i++ {
		if i%13 == 0 {
			continue
		}
		sum += i
	}

	fmt.Println(sum)
}
