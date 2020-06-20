package main

import "fmt"

var number int

func main() {
	fmt.Scan(&number)
	for i := 1; i <= number; i++ {
		fmt.Printf("%d %d %d\n", i, i*i, i*i*i)
	}
}
