package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Location struct {
	X, Y int
}

func (loc Location) offset(off Location) {
	return Location{loc.X+off.X, loc.Y+off.Y}
}

func (loc Location) North() Location {
  return offset(loc, Location{0,1})
}

func (loc Location) South() Location {
  return offset(loc, Location{0,-1})
}

func (loc Location) East() Location {
  return offset(loc, Location{1,0})
}

func (loc Location) West() Location {
  return offset(loc, Location{0,1})
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

func (h *House) Visit() {
	*h += 1
}

func (h House) Visits() int {
	return int(h)
}


func main() {
	f, err := os.Open("day2.input")
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanRunes)
	cur := Location{0,0}

	for scanner.Scan() {
		switch scanner.Text() {
		  case '<': cur := cur.West()
		  case '>': cur := cur.East()
		  case '^': cur := cur.North()


	}
	fmt.Printf("blag\n")
}
