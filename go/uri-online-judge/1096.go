package main

import "fmt"

func main() {
	for i := 1; i <= 9; i += 2 {
		fmt.Printf("I=%d J=%d\n", i, 7)
		fmt.Printf("I=%d J=%d\n", i, 6)
		fmt.Printf("I=%d J=%d\n", i, 5)
	}
}
