package main

import "testing"

func TestGreeting(t *testing.T) {
	got := greeting("TEST")
	want := "<b>TEST</b>"

	if got != want {
		t.Errorf("greeting('TEST') \n got: %v \n want: %v", got, want)
	}
}
