package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Coord struct {
	X, Y int
}

func NewCoord(s string) Coord {
	p := strings.Split(s, ",")
	x, err := strconv.Atoi(p[0])
	if err != nil {
		panic(fmt.Sprintf("Error converting to int: %v", err))
	}
	y, err := strconv.Atoi(p[1])
	if err != nil {
		panic(fmt.Sprintf("Error converting to int: %v", err))
	}
	return Coord{x, y}
}

type CoordPair struct {
	A, B Coord
}

func NewCoordPair(s0, s1 string) *CoordPair {
	return &CoordPair{NewCoord(s0), NewCoord(s1)}
}

var cmdMap map[string]func(f *Field) Coord

type Field map[Coord]struct{}

func (f *Field) turnon(c Coord) {
	(*f)[c] = struct{}{}
}

func (f *Field) turnoff(c Coord) {
	delete(*f, c)
}

func (f *Field) toggle(c Coord) {
	if _, ok := (*f)[c]; ok {
		f.turnoff(c)
	} else {
		f.turnon(c)
	}
}

func init() {
	cmdMap = make(map[string]func(f *Field) Coord, 3)
	cmdMap["turnon"] = func() { f.turnon() }
	cmdMap["turnoff"] = turnoff
	cmdMap["toggle"] = toggle
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
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

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
	parts := strings.Split(s, " ")
	if len(parts) != 3 {
		return errors.New(fmt.Sprintf("Malformed line: %s", s))
	}

	cmd, ok := cmdMap[parts[0]]
	if !ok {
		return errors.New(fmt.Sprintf("unknown command %s", parts[0]))
	}

	r := NewCoordPair{parts[1], parts[2]}

	return nil
}
