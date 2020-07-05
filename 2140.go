package main

import "fmt"

func main() {
	var n, m int

	notes := []int{100, 50, 20, 10, 5, 2}

	for true {
		fmt.Scan(&n, &m)

		if n+m == 0 {
			break
		}

		change := m - n
		answer := false

		for i := 0; i < 6 && !answer; i++ {
			for j := i; j < 6; j++ {
				if notes[i]+notes[j] == change {
					answer = true
					break
				}
			}
		}

		if answer {
			fmt.Println("possible")
		} else {
			fmt.Println("impossible")
		}
	}
}
