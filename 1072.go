package main

import "fmt"

var n, number, in, out int

func main() {

	fmt.Scan(&n)

	for n > 0 {
		fmt.Scan(&number)

		if number < 10 || number > 20 {
			out++
		} else {
			in++
		}

		n--
	}
	fmt.Printf("%d in\n", in)
	fmt.Printf("%d out\n", out)
}
