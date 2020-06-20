package main

import "fmt"

var acumulator, interval int

func main() {
	fmt.Scan(&acumulator)
	for interval <= 0 {
		fmt.Scan(&interval)
	}
	gaussSum := (interval * (interval - 1)) >> 1
	fmt.Println(acumulator*interval + gaussSum)
}
