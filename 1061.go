package main

import "fmt"

type date struct {
	text    string
	day     int
	hours   int
	minutes int
	seconds int
}

func (ob date) timeInSeconds() int {
	return (60 * 60 * 24 * ob.day) + (3600 * ob.hours) + (60 * ob.minutes) + (1 * ob.seconds)
}

func secondsToHMS(seconds int) (int, int, int, int) {
	days := (seconds / (60 * 60 * 24))
	hours := (seconds % (60 * 60 * 24)) / 3600
	minutes := (seconds % 3600) / 60
	seconds = (seconds % 60)

	return days, hours, minutes, seconds
}

func main() {
	time := [2]date{}

	for i := range time {
		fmt.Scanf("%v %d", &time[i].text, &time[i].day)
		fmt.Scanf("%d : %d : %d", &time[i].hours, &time[i].minutes, &time[i].seconds)
	}

	secondsBegin := time[0].timeInSeconds()
	secondsEnd := time[1].timeInSeconds()

	difference := secondsEnd - secondsBegin

	if difference < 0 {
		difference = (60 * 60 * 24) - (-1 * difference)
	}

	days, hours, minutes, seconds := secondsToHMS(difference)
	fmt.Printf("%d dia(s)\n%d hora(s)\n%d minuto(s)\n%d segundo(s)\n", days, hours, minutes, seconds)
}
