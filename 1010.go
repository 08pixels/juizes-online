package main

import "fmt"

var (
	idA       int
	quantityA float64
	priceA    float64
)

var (
	idB       int
	quantityB float64
	priceB    float64
)

func main() {
	fmt.Scanf("%d %f %f", &idA, &quantityA, &priceA)
	fmt.Scanf("%d %f %f", &idB, &quantityB, &priceB)

	answer := quantityA*priceA + quantityB*priceB
	fmt.Printf("VALOR A PAGAR: R$ %.2f\n", answer)
}
