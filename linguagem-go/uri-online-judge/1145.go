package main

import "fmt"

var x, y int

func main() {

	fmt.Scan(&x, &y)

	for i := 0; i < y; i += x {
		for j := 1; j < x; j++ {
			fmt.Printf("%d ", i+j)
		}
		fmt.Println(i + x)
	}
}
