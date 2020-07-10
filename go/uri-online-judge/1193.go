package main

import (
	"fmt"
	"strconv"
)

func main() {
	var cases, b int
	var number, base string

	fmt.Scan(&cases)

	for i := 1; i <= cases; i++ {
		fmt.Scan(&number, &base)

		if base == "bin" {
			b = 2
		} else if base == "dec" {
			b = 10
		} else {
			b = 16
		}

		n, _ := strconv.ParseInt(number, b, 32)

		fmt.Printf("Case %d:\n", i)

		if base != "dec" {
			fmt.Printf("%d dec\n", n)
		}

		if base != "hex" {
			fmt.Printf("%x hex\n", n)
		}

		if base != "bin" {
			fmt.Printf("%b bin\n", n)
		}
		fmt.Printf("\n")
	}
}
