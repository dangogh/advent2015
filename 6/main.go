package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Coord struct {
	X, Y int
}

type Section struct {
	A, B Coord
}

// normalizes bounds of box -- lower left to upper right
func NewSection(a, b Coord) Section {
	x0, x1 := a.X, b.X
	y0, y1 := a.Y, b.Y
	if x0 > x1 {
		x0, x1 = x1, x0
	}
	if y0 > y1 {
		y0, y1 = y1, y0
	}
	return Section{Coord{x0, y0}, Coord{x1, y1}}
}

type Grid map[Coord]struct{}

var theGrid Grid

func (g Grid) apply(s Section, f func(Coord)) {
	for i := s.A.X; i <= s.B.X; i++ {
		for j := s.A.Y; j <= s.B.Y; j++ {
			f(Coord{i, j})
		}
	}
}

func (g Grid) turnonSection(s Section) {
	g.apply(s, func(c Coord) {
		g.turnonCoord(c)
	})
}

func (g Grid) turnoffSection(s Section) {
	g.apply(s, func(c Coord) {
		g.turnoffCoord(c)
	})
}

func (g Grid) toggleSection(s Section) {
	g.apply(s, func(c Coord) {
		g.toggleCoord(c)
	})
}

func (g Grid) turnonCoord(c Coord) {
	g[c] = struct{}{}
}

func (g Grid) turnoffCoord(c Coord) {
	delete(g, c)
}

func (g Grid) toggleCoord(c Coord) {
	if _, ok := g[c]; ok {
		g.turnoffCoord(c)
	} else {
		g.turnonCoord(c)
	}
}

var cmdMap map[string]func(Section)

func init() {
	theGrid = make(map[Coord]struct{})
	cmdMap = map[string]func(s Section){
		"turn on": func(s Section) {
			theGrid.apply(s, func(c Coord) { theGrid.turnonCoord(c) })
		},
		"turn off": func(s Section) {
			theGrid.apply(s, func(c Coord) { theGrid.turnoffCoord(c) })
		},
		"toggle": func(s Section) {
			theGrid.apply(s, func(c Coord) { theGrid.toggleSection(s) })
		},
	}
}

func main() {
	var reader *bufio.Reader
	if len(os.Args) == 1 {
		reader = bufio.NewReader(os.Stdin)
	} else {
		f, err := os.Open(os.Args[1])
		defer f.Close()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Reading from %s\n", os.Args[1])
		reader = bufio.NewReader(f)
	}
	scanner := bufio.NewScanner(reader)

	//display := make(Display)
	for scanner.Scan() {
		line := scanner.Text()
		err := parseAndExec(line)
		if err != nil {
			log.Fatal(err)
		}
	}

}

func parseAndExec(s string) error {
	var cmd, loc string

	for c, _ := range cmdMap {
		if strings.HasPrefix(s, c) {
			cmd = c
			loc = strings.TrimSpace(strings.TrimPrefix(s, cmd+" "))
			break
		}
	}
	// make sure one was found..
	if cmd == "" {
		return fmt.Errorf("Unknown command %s", s)
	}

	var x0, y0, x1, y1 int

	fmt.Sscanf(loc, "%d,%d through %d,%d", &x0, &y0, &x1, &y1)
	//fmt.Printf("cmd: %s, %s\n", cmd, loc)
	//fmt.Printf(" %d,%d -> %d,%d\n", x0, y0, x1, y1)
	sec := Section{Coord{x0, y0}, Coord{x1, y1}}

	fmt.Printf("%s section %v\n", cmd, sec)
	cmdMap[cmd](sec)
	return nil
}
