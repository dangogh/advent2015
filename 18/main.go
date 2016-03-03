package main

import (
	"bufio"
	"fmt"
	"os"
)

/*var desc = `A light which is on stays on when 2 or 3 neighbors are on, and turns off otherwise.
A light which is off turns on if exactly 3 neighbors are on, and stays off otherwise.
`*/

type cell struct {
	x, y int
	on   bool
}

type grid [][]cell

func bounds(min, max, a, b int) (int, int) {
	if a < min {
		a = min
	}
	if b > max {
		b = max
	}
	return a, b
}

func (g grid) liveneighbors(c cell) int {
	lowerx, upperx := bounds(0, len(g)-1, c.x-1, c.x+1)
	lowery, uppery := bounds(0, len(g)-1, c.y-1, c.y+1)

	d := false
	if d {
		fmt.Printf("start  %d,%d -> %d,%d\n", lowerx, lowery, upperx, uppery)
		fmt.Printf("%s\n", g.String())
	}
	var n int

	if d {
		fmt.Printf("-------------------\n")
		fmt.Printf("  %d,%d -> %d,%d\n", lowery, lowerx, uppery, upperx)
	}

	for i := lowery; i <= uppery; i++ {
		for j := lowerx; j <= upperx; j++ {
			if j == c.x && i == c.y {
				continue
			}
			if d {
				fmt.Printf("%d,%d  %v -- %v\n", i, j, c, g[j][i])
			}
			if g[j][i].on {
				n++
			}
		}
	}
	return n
}

func (g grid) livecells() int {
	var live int
	for _, r := range g {
		for _, c := range r {
			if c.on {
				live++
			}
		}
	}
	return live
}

func (g grid) willLive(c cell) bool {
	// if a corner light,  always on
	if (c.x == 0 || c.x == len(g)-1) && (c.y == 0 || c.y == len(g)-1) {
		return true
	}
	switch g.liveneighbors(c) {
	case 2:
		// don't change
		return c.on
	case 3:
		return true
	default:
		return false
	}
}

func (g grid) nextStage() grid {
	newg := make(grid, 0, len(g))
	for _, row := range g {
		var newrow []cell
		for _, c := range row {
			newc := cell{c.x, c.y, g.willLive(c)}
			newrow = append(newrow, newc)
		}
		newg = append(newg, newrow)
	}
	return newg
}

func (g grid) String() string {
	var s string
	for _, row := range g {
		for _, c := range row {
			if c.on {
				s += "#"
			} else {
				s += "."
			}
		}
		s += "\n"
	}
	return s
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: ... <input>\n")
		os.Exit(1)
	}
	f, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	s := bufio.NewScanner(f)
	s.Split(bufio.ScanLines)
	var j int
	var g grid
	for s.Scan() {
		t := s.Text()
		row := make([]cell, 0, len(t))
		for i, st := range t {
			var on bool
			if st == rune('#') {
				on = true
			}
			row = append(row, cell{x: j, y: i, on: on})
		}
		j++
		g = append(g, row)
	}
	g[0][0].on = true
	g[0][len(g)-1].on = true
	g[len(g)-1][0].on = true
	g[len(g)-1][len(g)-1].on = true

	fmt.Println("Initial state:\n", g.String())
	fmt.Printf("Start with %d live cells\n", g.livecells())
	for i := 0; i < 100; i++ {
		g = g.nextStage()
		fmt.Println(g.String())
		fmt.Println(g.livecells())
	}
}
