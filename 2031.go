package main

import "fmt"

func main() {
	var cases int
	var player1, player2 string
	var answer string

	fmt.Scan(&cases)

	for cases != 0 {
		fmt.Scan(&player1, &player2)

		answer = "Jogador 1 venceu"

		if player1 == player2 {
			if player1 == "papel" {
				answer = "Ambos venceram"
			} else if player1 == "pedra" {
				answer = "Sem ganhador"
			} else {
				answer = "Aniquilacao mutua"
			}

		} else if player2 == "ataque" {
			answer = "Jogador 2 venceu"
		} else if player1 == "papel" {
			answer = "Jogador 2 venceu"
		}

		fmt.Println(answer)
		cases--
	}
}
