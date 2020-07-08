package main

import "fmt"

var number int

func main() {
	for true {
		fmt.Scan(&number)

		if number == 0 {
			break
		}

		if number&1 == 1 {
			number++
		}
		fmt.Println(number*5 + 20)
	}
}
