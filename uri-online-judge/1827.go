package main

import "fmt"

var number int
var oneBegin, oneEnd, four int

func main() {
	for true {
		_, err := fmt.Scan(&number)

		if err != nil {
			break
		}

		oneBegin = int(number / 3)
		oneEnd = number - 1 - oneBegin
		four = int(number / 2)

		for i := 0; i < number; i++ {
			row := ""
			for j := 0; j < number; j++ {
				if i == j && i == four {
					row = row + "4"
				} else if (i >= oneBegin && i <= oneEnd) && (j >= oneBegin && j <= oneEnd) {
					row = row + "1"
				} else if i == j {
					row = row + "2"
				} else if j == number-1-i {
					row = row + "3"
				} else {
					row = row + "0"
				}
			}
			fmt.Println(row)
		}
		fmt.Printf("\n")
	}
}
