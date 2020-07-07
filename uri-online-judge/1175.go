package main

import "fmt"

var arr = [20]int{}

func main() {

	for i := 1; i <= 20; i++ {
		fmt.Scan(&arr[20-i])
	}

	for i := 0; i < 20; i++ {
		fmt.Printf("N[%d] = %d\n", i, arr[i])
	}
}
