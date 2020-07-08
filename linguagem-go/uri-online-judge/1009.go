package main

import "fmt"

var name string
var salary float64
var sold float64

func main() {
	fmt.Scanf("%v\n%f\n%f", &name, &salary, &sold)
	fmt.Printf("TOTAL = R$ %.2f\n", salary+0.15*sold)
}
