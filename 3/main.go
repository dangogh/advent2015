package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	//"sort"
	//"strconv"
	//"strings"
)

type Location struct {
	X, Y int
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
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for scanner.ScanBytes() {
		line := scanner.Text()
		if err != nil {
			log.Fatal(err)
		}

	}
	fmt.Printf("blag\n")
}
