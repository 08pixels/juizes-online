package main

import "fmt"

var cases int
var word string

func main() {
	fmt.Scan(&cases)

	for cases != 0 {
		fmt.Scan(&word)

		if len(word) == 5 {
			fmt.Println(3)
		} else {
			if (word[0] == 'o' && word[1] == 'n') || (word[0] == 'o' && word[2] == 'e') || (word[1] == 'n' && word[2] == 'e') {
				fmt.Println(1)
			} else {
				fmt.Println(2)
			}
		}
		cases--
	}
}
