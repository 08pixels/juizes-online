package main

import "fmt"

var banknotes = []int{
	100, 50, 20, 10, 5, 2, 1,
}

var value int

func main() {
	fmt.Scanf("%d", &value)
	fmt.Println(value)
	for _, note := range banknotes {
		fmt.Printf("%d nota(s) de R$ %d,00\n", value/note, note)
		value %= note
	}
}
