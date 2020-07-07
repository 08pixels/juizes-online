package main

import "fmt"

var n, number int

func main() {
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&number)

		if number != 0 {
			if number%2 == 0 {
				fmt.Printf("EVEN ")
			} else {
				fmt.Printf("ODD ")
			}
		}

		if number < 0 {
			fmt.Println("NEGATIVE")
		} else if number > 0 {
			fmt.Println("POSITIVE")
		} else {
			fmt.Println("NULL")
		}
	}
}
