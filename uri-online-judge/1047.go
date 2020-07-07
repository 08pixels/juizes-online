package main

import "fmt"

var beginHour, endHour int
var beginMinute, endMinute int

func main() {
	fmt.Scan(&beginHour, &beginMinute, &endHour, &endMinute)

	if endHour < beginHour || ((endHour == beginHour) && (beginMinute >= endMinute)) {
		endHour += 24
	}

	beginInMinutes := 60*beginHour + beginMinute
	endInMinutes := 60*endHour + endMinute

	total := endInMinutes - beginInMinutes
	hours := total / 60
	minutes := total % 60

	fmt.Printf("O JOGO DUROU %d HORA(S) E %d MINUTO(S)\n", hours, minutes)

}
