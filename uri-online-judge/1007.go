package main

import "fmt"

var valueA int
var valueB int
var valueC int
var valueD int

func main() {
	fmt.Scanf("%d\n%d\n%d\n%d", &valueA, &valueB, &valueC, &valueD)
	fmt.Printf("DIFERENCA = %d\n", valueA*valueB-valueC*valueD)
}
