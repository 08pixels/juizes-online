package main

import (
	"bufio"
	"fmt"
	"os"
)

var cases int
var buffer []byte

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buffer, 1005)
	fmt.Scan(&cases)

	for cases != 0 {
		err := scanner.Scan()

		if err == false {
			break
		}

		message := scanner.Text()
		messageRunes := []rune(message)

		for i := 0; i < len(messageRunes); i++ {
			if messageRunes[i] >= 'a' && messageRunes[i] <= 'z' {
				messageRunes[i] = messageRunes[i] + 3
			} else if messageRunes[i] >= 'A' && messageRunes[i] <= 'Z' {
				messageRunes[i] = messageRunes[i] + 3
			}
		}

		size := len(messageRunes) - 1
		for i := 0; i <= size/2; i++ {
			messageRunes[i], messageRunes[size-i] = messageRunes[size-i], messageRunes[i]-1
		}

		fmt.Println(string(messageRunes))
		cases--
	}
}
