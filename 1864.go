package main

import "fmt"

var number int

const phrase = "LIFE IS NOT A PROBLEM TO BE SOLVED"

func main() {
	fmt.Scan(&number)
	fmt.Println(phrase[:number])
}
