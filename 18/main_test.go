package main

import (
	"testing"
)

var tests = map[string]grid{
	"none": grid{
		[]cell{cell{x: 0, y: 0, on: false}},
	},
	"alloff": grid{
		[]cell{cell{0, 0, false}, cell{1, 0, false}, cell{2, 0, false}},
		[]cell{cell{0, 1, false}, cell{1, 1, false}, cell{2, 1, false}},
		[]cell{cell{0, 2, false}, cell{1, 2, false}, cell{2, 2, false}},
	},
	"allon": grid{
		[]cell{cell{0, 0, true}, cell{1, 0, true}, cell{2, 0, true}},
		[]cell{cell{0, 1, true}, cell{1, 1, true}, cell{2, 1, true}},
		[]cell{cell{0, 2, true}, cell{1, 2, true}, cell{2, 2, true}},
	},
}

func TestNoNeighbors(t *testing.T) {
	g := tests["none"]
	if g.liveneighbors(g[0][0]) != 0 {
		t.Errorf("lone cell has no neighbors")
	}
}

func TestAllOff(t *testing.T) {
	g := tests["alloff"]
	if g.liveneighbors(g[1][1]) != 0 {
		t.Errorf("no on neigbors")
	}
	// test corner/edge
	if g.liveneighbors(g[0][2]) != 0 {
		t.Errorf("no on neighbors from corner")
	}
	if g.liveneighbors(g[1][2]) != 0 {
		t.Errorf("no on neighbors from edge")
	}
}

func TestAllOn(t *testing.T) {
	g := tests["allon"]
	if g.liveneighbors(g[1][1]) != 8 {
		t.Errorf("all on neigbors")
	}
	// test corner
	if g.liveneighbors(g[2][0]) != 3 {
		t.Errorf("all on at corner")
	}
	// test edge
	if g.liveneighbors(g[2][1]) != 5 {
		t.Errorf("all on at edge")
	}
}
