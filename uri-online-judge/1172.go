package main

import "fmt"

var x int

func main() {
	for i := 0; i < 10; i++ {
		fmt.Scan(&x)
		if x <= 0 {
			x = 1
		}
		fmt.Printf("X[%d] = %d\n", i, x)
	}
}
