package main

import "fmt"

var x int
var y int

func main() {
	fmt.Scanf("%v\n%v", &x, &y)
	fmt.Printf("X = %v\n", x+y)
}
