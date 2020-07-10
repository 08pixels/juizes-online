package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		sentence := scanner.Text()
		dancing := []rune(sentence)
		isUpper := true

		for i := 0; i < len(dancing); i++ {
			if dancing[i] >= 'a' && dancing[i] <= 'z' {
				if isUpper {
					dancing[i] = 'A' + (dancing[i] - 'a')
				}

				isUpper = !isUpper
			} else if dancing[i] >= 'A' && dancing[i] <= 'Z' {
				if !isUpper {
					dancing[i] = 'a' + (dancing[i] - 'A')
				}

				isUpper = !isUpper
			}
		}
		fmt.Println(string(dancing))
	}

}
