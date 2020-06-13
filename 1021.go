package main

import "fmt"

var integer int
var fractional int

var noteBank = []int{100, 50, 20, 10, 5, 2}
var coinBank = []int{100, 50, 25, 10, 5, 1}

func main() {
	fmt.Scanf("%d.%d", &integer, &fractional)

	fmt.Println("NOTAS:")
	for i := 0; i < 6; i++ {
		fmt.Printf("%d nota(s) de R$ %d.00\n", integer/noteBank[i], noteBank[i])
		integer %= noteBank[i]
	}

	fractional += (100 * integer)
	fmt.Println("MOEDAS:")
	for i := 0; i < 6; i++ {
		fmt.Printf("%d moeda(s) de R$ %.2f\n", fractional/coinBank[i], float64(coinBank[i])/100.0)
		fractional %= coinBank[i]
	}
}
