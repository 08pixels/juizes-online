package main

import "fmt"

var limit int

func main() {
	fmt.Scan(&limit)

	for i := 1; i <= limit; i += 2 {
		fmt.Println(i)
	}
}
