package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

type Location struct {
	X, Y int
}

func (loc Location) offset(off Location) Location {
	return Location{loc.X + off.X, loc.Y + off.Y}
}

func (loc Location) North() Location {
	return loc.offset(Location{0, 1})
}

func (loc Location) South() Location {
	return loc.offset(Location{0, -1})
}

func (loc Location) East() Location {
	return loc.offset(Location{1, 0})
}

func (loc Location) West() Location {
	return loc.offset(Location{-1, 0})
}

type House int
type Neighborhood map[Location]*House

func (n Neighborhood) HouseAt(loc Location) House {
	return *(n[loc])
}

func (n Neighborhood) Visit(loc Location) {
	if _, ok := n[loc]; !ok {
		// hasn't yet been visited
		n[loc] = new(House)
	}
	n[loc].Visit()
}

func (n Neighborhood) String() string {
	s := ""
	for loc, visits := range n {
		s += fmt.Sprintf("[ %d, %d ]: %d\n", loc.X, loc.Y, *visits)
	}
	return s
}

func (h *House) Visit() {
	*h += 1
}

func (h House) Visits() int {
	return int(h)
}

func part1(n Neighborhood, reader *bufio.Reader) {
	cur := Location{0, 0}
	for {
		n.Visit(cur)
		r, _, err := reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(fmt.Sprintf("bad rune %s(%d)", string(r), int(r)))
		}
		switch r {
		case '>':
			cur = cur.East()
		case '<':
			cur = cur.West()
		case '^':
			cur = cur.North()
		case 'v':
			cur = cur.South()
		default:
			panic(fmt.Sprintf("Unknown direction %s", string(r)))
		}
	}
}

func part2(n Neighborhood, reader *bufio.Reader) {

	cur := make([]Location, 2)
	idx := 0
	for {
		n.Visit(cur[idx])
		r, _, err := reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(fmt.Sprintf("bad rune %s(%d)", string(r), int(r)))
		}
		switch r {
		case '>':
			cur[idx] = cur[idx].East()
		case '<':
			cur[idx] = cur[idx].West()
		case '^':
			cur[idx] = cur[idx].North()
		case 'v':
			cur[idx] = cur[idx].South()
		default:
			panic(fmt.Sprintf("Unknown direction %s", string(r)))
		}
		idx = 1 - idx
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
		reader = bufio.NewReader(f)
	}
	n := make(Neighborhood)

	//part1(n, reader)
	part2(n, reader)

	fmt.Printf("Neighborhood looks like this: %v\n", n)

	mvisits := 0
	for _, pv := range n {
		if *pv > 1 {
			mvisits++
		}
	}
	fmt.Printf("Number of houses with at least one visit: %d\n", len(n))
	fmt.Printf("Number of houses with multiple visits: %d\n", mvisits)
}
