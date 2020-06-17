package main

import "fmt"

var limit int

func main() {
	fmt.Scan(&limit)

	if limit%2 == 0 {
		limit++
	}

	for i := 0; i < 6; i++ {
		fmt.Println(limit)
		limit += 2
	}
}
