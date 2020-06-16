package main

import (
	"fmt"
	"sort"
)

var a, b, c float64

func main() {

	fmt.Scan(&a, &b, &c)

	sides := []float64{a, b, c}
	sort.Float64s(sides)
	a, b, c = sides[2], sides[1], sides[0]
	sqrA := a * a
	sqrB := b * b
	sqrC := c * c

	if a >= b+c {
		fmt.Println("NAO FORMA TRIANGULO")
	} else {
		if sqrA == sqrB+sqrC {
			fmt.Println("TRIANGULO RETANGULO")
		}
		if sqrA > sqrB+sqrC {
			fmt.Println("TRIANGULO OBTUSANGULO")
		}
		if sqrA < sqrB+sqrC {
			fmt.Println("TRIANGULO ACUTANGULO")
		}

		if a == b && b == c {
			fmt.Println("TRIANGULO EQUILATERO")
		} else if a == b || b == c {
			fmt.Println("TRIANGULO ISOSCELES")
		}
	}
}
