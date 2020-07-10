package main

import "fmt"

var seconds int

func main() {
	fmt.Scanf("%d", &seconds)
	hours := seconds / 3600
	seconds %= 3600
	minutes := seconds / 60
	fmt.Printf("%d:%d:%d\n", hours, minutes, seconds%60)
}
