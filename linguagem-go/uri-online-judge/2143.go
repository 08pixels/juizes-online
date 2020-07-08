package main

import "fmt"

func main() {
	var cases, number int

	for true {
		fmt.Scan(&cases)

		if cases == 0 {
			break
		}

		for i := 0; i < cases; i++ {
			fmt.Scan(&number)

			if number&1 == 0 {
				fmt.Println((number-2)*2 + 2)
			} else {
				fmt.Println((number-1)*2 + 1)
			}
		}
	}
}
