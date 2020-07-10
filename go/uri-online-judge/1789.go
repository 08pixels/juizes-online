package main

import "fmt"

var number, velocity int

func main() {
	for true {
		_, err := fmt.Scan(&number)
		if err != nil {
			break
		}
		answer := 0
		for i := 0; i < number; i++ {
			fmt.Scan(&velocity)

			if answer < velocity {
				answer = velocity
			}
		}

		if answer < 10 {
			fmt.Println(1)
		} else if answer < 20 {
			fmt.Println(2)
		} else {
			fmt.Println(3)
		}
	}
}
