package main

import "fmt"

var beginGame, endGame int

func main() {
	fmt.Scan(&beginGame, &endGame)

	if beginGame > endGame {
		fmt.Printf("O JOGO DUROU %d HORA(S)\n", 24-beginGame+endGame)
	} else if beginGame < endGame {
		fmt.Printf("O JOGO DUROU %d HORA(S)\n", endGame-beginGame)
	} else {
		fmt.Printf("O JOGO DUROU %d HORA(S)\n", 24)
	}
}
