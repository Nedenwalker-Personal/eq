package main

import "fmt"

type TimeFrame struct {
	Name string
	StartRealTimeIn24HourClock int
	EndRealTimeIn24HourClock int
}

const hoursInDay int = 24
const minutesInHour int = 60
const secondsInMinute int = 60

func Name(tf TimeFrame) string {
	return tf.Name
}

func DurationInSeconds(tf TimeFrame) int {
	durationInHours := hoursInDay - tf.StartRealTimeIn24HourClock + tf.EndRealTimeIn24HourClock
	return secondsInMinute * minutesInHour * durationInHours
}

func main() {
	var tf TimeFrame
	tf.Name = "Night"
	tf.StartRealTimeIn24HourClock = 20
	tf.EndRealTimeIn24HourClock = 6
	fmt.Printf("%+v\n", tf)

	tfd := DurationInSeconds(tf)

	for i := 0; i < tfd; i++ {
		tl := tfd - i
		if tl == 5 * 60 {
			fmt.Printf("Time is: %d", i)
		}
	}
}
