package main

import "fmt"

var number int
var value float32
var hours float32

func main() {
	fmt.Scanf("%d\n%f\n%f", &number, &value, &hours)
	fmt.Printf("NUMBER = %d\n", number)
	fmt.Printf("SALARY = U$ %.2f\n", value*hours)
}
