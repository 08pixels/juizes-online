package main

import "fmt"

var cases int
var amount int
var specie string

var habbits, frogs, mouses int

func main() {

	fmt.Scan(&cases)
	for i := 0; i < cases; i++ {
		fmt.Scan(&amount, &specie)

		if specie == "S" {
			frogs += amount
		} else if specie == "C" {
			habbits += amount
		} else {
			mouses += amount
		}
	}

	total := mouses + frogs + habbits
	pHabbits := float64(100*habbits) / float64(total)
	pMouses := float64(100*mouses) / float64(total)
	pFrogs := float64(100*frogs) / float64(total)

	fmt.Printf("Total: %d cobaias\n", total)
	fmt.Printf("Total de coelhos: %d\nTotal de ratos: %d\nTotal de sapos: %d\n", habbits, mouses, frogs)
	fmt.Printf("Percentual de coelhos: %.2f %%\nPercentual de ratos: %.2f %%\nPercentual de sapos: %.2f %%\n", pHabbits, pMouses, pFrogs)
}
