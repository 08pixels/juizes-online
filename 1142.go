package main

import "fmt"

var number int

func main() {
	fmt.Scan(&number)

	for i := 0; i < number; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Printf("%d ", 4*i+j)
		}
		fmt.Println("PUM")
	}
}
