package main

import "fmt"

var number int
var counter int

func main() {
	fmt.Scan(&number)

	for i := 0; i < 1000; i++ {
		fmt.Printf("N[%d] = %d\n", i, counter)
		counter = (counter + 1) % number
	}
}
