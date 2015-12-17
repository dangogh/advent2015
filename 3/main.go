package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Location struct {
	X, Y int
}

type House int
type Neighborhood map[Location]*House

func (n *Neighborhood) Visit(l Location) {
	if h, ok := n[l]; !ok {
		// hasn't yet been visited
		n[l] = &House{0}
	}
	n[l].Visit()

}

func (h *House) Visit() {
	h += 1
}

func (h House) Visits() int {
	return int(h)
}

func main() {

	f, err := os.Open("day2.input")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	wrappingTotal := 0
	ribbonTotal := 0
	for scanner.Scan() {
		line := scanner.Text()
		b, err := NewBox(line)
		if err != nil {
			log.Fatal(err)
		}

		wrappingTotal += b.WrappingArea()
		p := b.Perimeters()
		ribbonTotal += p[0] + b.Volume()
	}
	fmt.Printf("Total wrapping paper area of %d sq ft\n", wrappingTotal)
	fmt.Printf("Total ribbon length of %d ft\n", ribbonTotal)
}
