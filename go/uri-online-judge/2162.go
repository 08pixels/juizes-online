package main

import "fmt"

func main() {
	var cases int

	fmt.Scan(&cases)
	numbers := make([]int, cases)

	for i := 0; i < cases; i++ {
		fmt.Scan(&numbers[i])
	}

	peak := 0
	answer := true

	for i := 1; i < cases && answer; i++ {
		last := numbers[i-1]

		if last < numbers[i] {
			if peak == 1 {
				answer = false
			}
			peak = 1
		} else if last > numbers[i] {
			if peak == -1 {
				answer = false
			}
			peak = -1
		} else {
			answer = false
		}
	}

	if answer {
		fmt.Println("1")
	} else {
		fmt.Println("0")
	}
}
