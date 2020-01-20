package main

import "testing"

func TestName(t *testing.T) {
	var tf TimeFrame
	tf.Name = "Night"
	got := Name(tf)
	want := "Night"
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func TestDurationInSeconds(t *testing.T) {
	var tf TimeFrame
	tf.StartRealTimeIn24HourClock = 20
	tf.EndRealTimeIn24HourClock = 6
	got := DurationInSeconds(tf)
	want := 36000
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
