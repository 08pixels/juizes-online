package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		form := scanner.Text()

		if len(form) <= 80 {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
