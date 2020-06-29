package main

import "fmt"

var a, b, c, d int

func main() {
	fmt.Scan(&a, &b, &c, &d)
	fmt.Println(a + b + c + d - 3)
}
