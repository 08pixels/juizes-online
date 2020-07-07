package main

import "fmt"

func main() {
	var p, j1, j2, r, a int
	fmt.Scan(&p, &j1, &j2, &r, &a)

	winner := 1

	if p == 0 && (j1&1 == j2&1) {
		winner = 2
	} else if p == 1 && (j1&1 != j2&1) {
		winner = 2
	}

	if r == 1 && a == 1 {
		winner = 2
	} else if r == 0 && a == 1 {
		winner = 1
	} else if r == 1 && a == 0 {
		winner = 1
	}

	fmt.Printf("Jogador %d ganha!\n", winner)
}
