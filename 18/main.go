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

func (g grid) liveneighbors(c cell) int {
	lowerx, upperx := c.x-1, c.x+1
	lowery, uppery := c.y-1, c.y+1
	if lowerx == -1 {
		lowerx = 0
	}
	if lowery == -1 {
		lowery = 0
	}
	if upperx == len(g) {
		upperx = len(g) - 1
	}
	if uppery == len(g[0]) {
		uppery = len(g[0]) - 1
	}
	var n int
	for j := lowerx; j <= upperx; j++ {
		for i := lowery; i <= uppery; i++ {
			if j == c.x && i == c.y {
				continue
			}
			if g[j][i].on {
				n++
			}
		}
	}
	return n
}

func (g grid) willLive(c cell) bool {
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
		g = append(g, row)
	}

	for range []int{1, 2, 3, 4, 5} {
		fmt.Println(g.String())
		g = g.nextStage()
	}
}
