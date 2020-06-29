package main

import (
	"bufio"
	"fmt"
	"os"
)

var cases int
var course string

func main() {
	fmt.Scan(&cases)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)

	for cases != 0 {
		scanner.Scan()
		cases--
	}
	fmt.Println("Ciencia da Computacao")
}
