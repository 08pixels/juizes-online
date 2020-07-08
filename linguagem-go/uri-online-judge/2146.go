package main

import "fmt"

func main() {
	var password int

	for true {
		_, err := fmt.Scan(&password)

		if err != nil {
			break
		}

		fmt.Println(password - 1)
	}
}
