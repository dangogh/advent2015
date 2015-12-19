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
type Grid map[Location]House

func (loc Location) North() Location {
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
