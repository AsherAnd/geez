package main

import "testing"

func TestNumToGeez(t *testing.T) {
	got := numToGeez(1)
	want := "፩"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
