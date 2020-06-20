package main

import "fmt"

var alcohol, gas, diesel, option int

func main() {

	for option != 4 {
		fmt.Scan(&option)

		if option == 1 {
			alcohol++
		} else if option == 2 {
			gas++
		} else if option == 3 {
			diesel++
		}
	}

	fmt.Println("MUITO OBRIGADO")
	fmt.Printf("Alcool: %d\n", alcohol)
	fmt.Printf("Gasolina: %d\n", gas)
	fmt.Printf("Diesel: %d\n", diesel)
}
