package main

import (
	"testing"
)

func TestNoNeighbors(t *testing.T) {
	c := cell{x: 0, y: 0, on: false}
	g := grid{[]cell{c}}
	if g.liveneighbors(c) != 0 {
		t.Fail()
	}
}
