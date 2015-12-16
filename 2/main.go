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

func dimensions(line string) ([]int, error) {
	p := make([]int, 0)
	for _, s := range strings.Split(line, "x") {
		v, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
		p = append(p, v)
	}
	if len(p) != 3 {
		return nil, errors.New(fmt.Sprintf("Expected 3 parts to %s but got %d", line, len(p)))
	}

	sort.Ints(p)
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

	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		dim, err := dimensions(line)
		if err != nil {
			log.Fatal(err)
		}

		a := 3*dim[0]*dim[1] + 2*dim[0]*dim[2] + 2*dim[1]*dim[2]
		total += a
	}
	fmt.Printf("Total of %d sq ft\n", total)
}
