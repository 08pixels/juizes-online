package main

import "fmt"

var odd = [5]int{}
var even = [5]int{}

var counterOdd, counterEven int
var number, count int

func printArray(name string, arr [5]int) {

	if name == "impar" {
		count = counterOdd
	} else {
		count = counterEven
	}

	for i := 0; i < count; i++ {
		fmt.Printf("%s[%d] = %d\n", name, i, arr[i])
	}
}

func main() {

	for i := 0; i < 15; i++ {
		fmt.Scan(&number)

		if number&1 == 1 {
			odd[counterOdd] = number
			counterOdd++
		} else {
			even[counterEven] = number
			counterEven++
		}

		if counterOdd == 5 {
			printArray("impar", odd)
			counterOdd = 0
		}
		if counterEven == 5 {
			printArray("par", even)
			counterEven = 0
		}
	}
	printArray("impar", odd)
	printArray("par", even)

}
