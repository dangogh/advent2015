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
	"alternating": grid{
		[]cell{cell{0, 0, true}, cell{1, 0, false}, cell{2, 0, true}},
		[]cell{cell{0, 1, false}, cell{1, 1, true}, cell{2, 1, false}},
		[]cell{cell{0, 2, true}, cell{1, 2, false}, cell{2, 2, true}},
	},
	"alternating2": grid{
		[]cell{cell{0, 0, false}, cell{1, 0, true}, cell{2, 0, false}},
		[]cell{cell{0, 1, true}, cell{1, 1, false}, cell{2, 1, true}},
		[]cell{cell{0, 2, false}, cell{1, 2, true}, cell{2, 2, false}},
	},
}

func TestGrid(t *testing.T) {
	for _, g := range tests {
		for j := range g {
			for i, c := range g[j] {
				if c.y != j || c.x != i {
					t.Errorf("Cell %v not at %d,%d", c, j, i)
				}
			}
		}
	}
}

func TestNoNeighbors(t *testing.T) {
	g := tests["none"]
	c := g[0][0]
	if g.liveneighbors(c) != 0 {
		t.Errorf("lone cell has no neighbors -- found %d\n", g.liveneighbors(c))
	}
	if g.willLive(c) {
		t.Errorf("lone cell turns off -- found %d\n", g.liveneighbors(c))
	}
}

func TestAllOff(t *testing.T) {
	g := tests["alloff"]
	c := g[1][1]
	if g.liveneighbors(c) != 0 {
		t.Errorf("no on neigbors -- found %d\n", g.liveneighbors(c))
	}
	if g.willLive(c) {
		t.Errorf("alloff stays off -- found %d\n", g.liveneighbors(c))
	}

	// test corner/edge
	c = g[0][2]
	if g.liveneighbors(c) != 0 {
		t.Errorf("no on neighbors from corner -- found %d\n", g.liveneighbors(c))
	}
	if g.willLive(c) {
		t.Errorf("alloff corner stays off -- found %d\n", g.liveneighbors(c))
	}

	c = g[1][2]
	if g.liveneighbors(c) != 0 {
		t.Errorf("no on neighbors from edge -- found %d\n", g.liveneighbors(c))
	}
	if g.willLive(c) {
		t.Errorf("alloff edge stays off -- found %d\n", g.liveneighbors(c))
	}
}

func TestAllOn(t *testing.T) {
	g := tests["allon"]
	c := g[1][1]
	if g.liveneighbors(c) != 8 {
		t.Errorf("allon neigbors -- found %d\n", g.liveneighbors(c))
	}
	if g.willLive(c) {
		t.Errorf("allon center turns off -- found %d\n", g.liveneighbors(c))
	}

	// test corner
	c = g[2][0]
	if g.liveneighbors(c) != 3 {
		t.Logf("GRID:\n%+v\n", g)
		t.Errorf("allon at corner -- found %d\n", g.liveneighbors(c))
	}
	if !g.willLive(c) {
		t.Errorf("allon corner stays on -- found %d\n", g.liveneighbors(c))
	}

	// test edge
	c = g[2][1]
	if g.liveneighbors(c) != 5 {
		t.Errorf("allon at edge -- found %d\n", g.liveneighbors(c))
	}
	if g.willLive(c) {
		t.Errorf("allon edge turns off -- found %d\n", g.liveneighbors(c))
	}
}

func TestAlternating(t *testing.T) {
	g := tests["alternating"]
	c := g[0][0]
	if g.liveneighbors(c) != 1 {
		t.Errorf("alt corner -- found %d\n", g.liveneighbors(c))
	}
	if g.willLive(c) {
		t.Errorf("alt corner turns off -- found %d\n", g.liveneighbors(c))
	}

	c = g[1][0]
	if g.liveneighbors(c) != 3 {
		t.Errorf("alt at edge -- found %d\n", g.liveneighbors(c))
	}
	if !g.willLive(c) {
		t.Errorf("alt edge turns on -- found %d\n", g.liveneighbors(c))
	}

	c = g[1][1]
	if g.liveneighbors(c) != 4 {
		t.Errorf("alt in center -- found %d\n", g.liveneighbors(c))
	}
	if g.willLive(c) {
		t.Errorf("alt in center turns off -- found %d\n", g.liveneighbors(c))
	}
}

func TestAlternating2(t *testing.T) {
	g := tests["alternating2"]
	c := g[0][0]
	if g.liveneighbors(c) != 2 {
		t.Errorf("alt2 corner -- found %d\n", g.liveneighbors(c))
	}
	if g.willLive(c) {
		t.Errorf("alt2 corner turns off -- found %d\n", g.liveneighbors(c))
	}

	c = g[1][0]
	if g.liveneighbors(c) != 2 {
		t.Errorf("alt2 edge -- found %d\n", g.liveneighbors(c))
	}
	if !g.willLive(c) {
		t.Errorf("alt2 edge stays on -- found %d\n", g.liveneighbors(c))
	}

	c = g[1][1]
	if g.liveneighbors(c) != 4 {
		t.Errorf("alt2 in center -- found %d\n", g.liveneighbors(c))
	}
	if g.willLive(c) {
		t.Errorf("alt2 edge turns off -- found %d\n", g.liveneighbors(c))
	}

	c = g[2][1]
	if g.liveneighbors(c) != 2 {
		t.Errorf("alt2 edge w/2 neighbors -- got %d with\n%s", g.liveneighbors(c), g.String())
	}
	if !g.willLive(c) {
		t.Errorf("edge comes alive -- found %d\n", g.liveneighbors(c))
	}
}
