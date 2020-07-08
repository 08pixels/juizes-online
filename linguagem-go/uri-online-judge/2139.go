package main

import "fmt"

func main() {
	var month, day int

	for true {
		_, err := fmt.Scan(&month, &day)

		if err != nil {
			break
		}

		christmas := 360
		currentDay := (month-1)*30 + day
		// 31, 29,31, 30, 31, 30,  31, 31,30,31,30
		// +1  0, +1, +1, +2, +2, +3, +4, +4,+5, +5

		switch month {
		case 2, 4, 5:
			currentDay++
		case 6, 7:
			currentDay += 2
		case 8:
			currentDay += 3
		case 9, 10:
			currentDay += 4
		case 11, 12:
			currentDay += 5
		}

		dif := christmas - currentDay
		switch {
		case dif == 1:
			fmt.Println("E vespera de natal!")
		case dif == 0:
			fmt.Println("E natal!")
		case dif < 0:
			fmt.Println("Ja passou!")
		default:
			fmt.Printf("Faltam %d dias para o natal!\n", dif)
		}
	}
}
