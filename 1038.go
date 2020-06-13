package main

import "fmt"

var price = []float32{4.00, 4.50, 5.0, 2.00, 1.50}
var quantity, product float32

func main() {
	fmt.Scan(&product, &quantity)
	fmt.Printf("Total: R$ %.2f\n", quantity*price[int(product)-1])
}
