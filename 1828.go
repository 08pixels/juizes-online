package main

import "fmt"

var sheldon, rajesh string
var cases int

var keyWins = map[string][]string{
	"pedra":   []string{"lagarto", "tesoura"},
	"papel":   []string{"pedra", "Spock"},
	"tesoura": []string{"papel", "lagarto"},
	"lagarto": []string{"Spock", "papel"},
	"Spock":   []string{"tesoura", "pedra"},
}

func main() {

	fmt.Scan(&cases)
	for i := 1; i <= cases; i++ {
		fmt.Scan(&sheldon, &rajesh)

		fmt.Printf("Caso #%d: ", i)
		if sheldon == rajesh {
			fmt.Println("De novo!")
		} else if rajesh == keyWins[sheldon][0] || rajesh == keyWins[sheldon][1] {
			fmt.Println("Bazinga!")
		} else {
			fmt.Println("Raj trapaceou!")
		}
	}
}
