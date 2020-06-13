package main

import "fmt"

var (
	a, b, c, d int
)

func main() {
	fmt.Scan(&a, &b, &c, &d)

	if c >= 0 && d >= 0 && b > c && d > a && (c+d > a+b) && (a%2 == 0) {
		fmt.Println("Valores aceitos")
	} else {
		fmt.Println("Valores nao aceitos")
	}
}
