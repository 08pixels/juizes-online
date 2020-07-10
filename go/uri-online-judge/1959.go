package main

import "fmt"

var side, length int64

func main() {
	fmt.Scan(&side, &length)
	fmt.Println(side * length)
}
