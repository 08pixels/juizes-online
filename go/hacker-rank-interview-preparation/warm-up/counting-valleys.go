package main

import "fmt"

func countingValleys(n int32, s string) int32 {
	last := int32(0)
	current := int32(0)
	answer := int32(0)

	for _, v := range s {
		last = current

		if v == 'D' {
			current--
		} else {
			current++
		}

		if current == -1 && last == 0 {
			answer++
		}
	}

	return answer
}

func main() {
	var number int32
	var walk string

	fmt.Scan(&number, &walk)
	fmt.Println(countingValleys(number, walk))
}
