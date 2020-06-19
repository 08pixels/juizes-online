package main

import "fmt"

func main() {
	j := 7
	for i := 1; i <= 9; i += 2 {
		fmt.Printf("I=%d J=%d\n", i, j)
		fmt.Printf("I=%d J=%d\n", i, j-1)
		fmt.Printf("I=%d J=%d\n", i, j-2)
		j += 2
	}
}
