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

type Box []int

func (b Box) Dimensions() []int {
	return []int(b)
}

func (b Box) Perimeters() []int {
	var s []int
	for i, d0 := range b.Dimensions() {
		for j, d1 := range b.Dimensions() {
			if i != j {
				s = append(s, 2*(d0+d1))
			}
		}
	}
	return s
}

func (b Box) SideAreas() []int {
	var s []int
	for i, d0 := range b.Dimensions() {
		for j, d1 := range b.Dimensions() {
			if i != j {
				s = append(s, d0*d1)
			}
		}
	}
	return s
}

func (b Box) SurfaceArea() int {
	s := b.SideAreas()
	t := 0
	for _, a := range s {
		t += a
	}
	return t
}

func (b Box) WrappingArea() int {
	s := b.SideAreas()
	return b.SurfaceArea() + s[0]
}

func (b Box) Volume() int {
	v := 1
	for _, d := range b.Dimensions() {
		v = v * d
	}
	return v
}

func NewBox(line string) (*Box, error) {
	p := &Box{}
	for _, s := range strings.Split(line, "x") {
		v, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
		*p = append(*p, v)
	}
	if len(*p) != 3 {
		return nil, errors.New(fmt.Sprintf("Expected 3 dimensions, not %d on %s", len(*p), line))
	}

	// sort dimensions increasing order
	sort.Ints(*p)
	return p, nil
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
