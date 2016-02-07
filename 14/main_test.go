package main

import (
	"testing"
)

func TestDistanceAt(t *testing.T) {
	r := reindeer{Name: "Izzy", Speed: 3, Duration: 4, Rest: 5}

	d := [][]int{{1, 3}, {10, 15}, {18, 24}}
	for _, p := range d {
		d := r.DistanceAt(p[0])
		if d != p[1] {
			t.Errorf("Expected , got %d", p[1], d)
		}
	}
}
