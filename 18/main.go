package main

import (
	"bufio"
	"fmt"
	"os"
)

var desc = `A light which is on stays on when 2 or 3 neighbors are on, and turns off otherwise.
A light which is off turns on if exactly 3 neighbors are on, and stays off otherwise.
`

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
			fmt.Printf("%d,%d %+v\n", j, i, g)
			if g[j][i].on {
				n++
			}
		}
	}
	return n
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
	for s.Scan() {
		t := s.Text()
		fmt.Printf("%s\n", t)
	}
}
