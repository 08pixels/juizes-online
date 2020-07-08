package main

import (
	"bufio"
	"fmt"
	"os"
)

var crow string
var number int

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		crow = scanner.Text()

		if crow != "caw caw" {
			for i := range crow {
				if crow[i] == '*' {
					number += (1 << uint32(2-i))
				}
			}
		} else {
			fmt.Println(number)
			number = 0
		}
	}
}
