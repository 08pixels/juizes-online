package main

import "fmt"

var number int

func main() {

	for true {
		_, err := fmt.Scan(&number)

		if err != nil {
			break
		}

		for i := 0; i < number; i++ {
			for j := 0; j < number; j++ {
				if (number - i - 1) == j {
					fmt.Printf("2")
				} else if i == j {
					fmt.Printf("1")
				} else {
					fmt.Printf("3")
				}
			}
			fmt.Printf("\n")
		}
	}
}
