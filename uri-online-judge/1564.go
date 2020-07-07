package main

import "fmt"

var number int

func main() {

	for true {
		_, err := fmt.Scan(&number)
		if err != nil {
			break
		}

		if number == 0 {
			fmt.Println("vai ter copa!")
		} else {
			fmt.Println("vai ter duas!")
		}
	}
}
