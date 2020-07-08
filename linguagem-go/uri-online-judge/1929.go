package main

import "fmt"

var a, b, c, d int

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func isValid(a int, b int, c int) bool {
	if abs(b-c) < a && a < b+c {
		return true
	}
	if abs(a-c) < b && b < a+c {
		return true
	}
	if abs(b-a) < c && c < b+a {
		return true
	}

	return false
}

func main() {
	fmt.Scan(&a, &b, &c, &d)

	if isValid(a, b, c) || isValid(a, b, d) || isValid(a, c, d) || isValid(b, c, d) {
		fmt.Println("S")
	} else {
		fmt.Println("N")
	}
}
