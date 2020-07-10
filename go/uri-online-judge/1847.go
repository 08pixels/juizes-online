package main

import "fmt"

var day1, day2, day3 int

func main() {
	fmt.Scan(&day1, &day2, &day3)

	if day1 > day2 && day2 <= day3 {
		fmt.Println(":)")
	} else if day1 < day2 && day2 < day3 && (day2-day1) <= (day3-day2) {
		fmt.Println(":)")
	} else if day1 > day2 && day2 > day3 && (day1-day2) > (day2-day3) {
		fmt.Println(":)")
	} else if day1 == day2 && day2 < day3 {
		fmt.Println(":)")
	} else {
		fmt.Println(":(")
	}
}
