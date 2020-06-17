package main

import "fmt"

var limit int

func main() {

	fmt.Scan(&limit)

	for i := 2; i <= limit; i += 2 {
		fmt.Printf("%d^2 = %d\n", i, i*i)
	}
}
