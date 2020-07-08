package main

import "fmt"

func main() {
	number := 1

	for number != 0 {
		fmt.Scan(&number)

		for i := 1; i < number; i++ {
			fmt.Printf("%d ", i)
		}

		if number != 0 {
			fmt.Println(number)
		}
	}
}
