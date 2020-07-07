package main

import "fmt"

func main() {
	var departure, duration, fuse int

	fmt.Scan(&departure, &duration, &fuse)
	if departure == 0 {
		departure = 24
	}
	fmt.Println((departure + duration + fuse) % 24)
}
