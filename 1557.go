package main

import "fmt"

var number int

func main() {
	for true {
		fmt.Scan(&number)

		if number == 0 {
			break
		}

		for i := 0; i < number; i++ {
			for j := 0; j < number; j++ {
				if j != (number - 1) {
					fmt.Printf("%d", 1<<(i+j))
				} else {
					fmt.Printf("%d\n", 1<<(i+j))
				}
			}
		}
	}
}
