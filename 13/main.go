package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

//Carol would gain 10 happiness units by sitting next to Eric.

type matrix map[string]map[string]int

func (m matrix) String() string {
	order := []string{}
	for n, _ := range m {
		order = append(order, n)
	}
	table := []string{"\t" + strings.Join(order, "\t")}

	for _, n := range order {
		row := []string{n}
		for _, o := range order {
			row = append(row, fmt.Sprintf("%5d", m[n][o]))
		}
		table = append(table, strings.Join(row, "\t"))
	}
	return strings.Join(table, "\n")
}

func loadWeights(r io.Reader) (matrix, error) {
	ss := make(matrix)
	re := regexp.MustCompile(`(\w+) would (\w+) (\d+) happiness units by sitting next to (\w+).`)

	s := bufio.NewScanner(r)
	s.Split(bufio.ScanLines)
	for s.Scan() {
		t := s.Text()

		fmt.Println(t)
		m := re.FindSubmatch([]byte(t))
		if m == nil {
			return nil, fmt.Errorf("no match on line: %s", t)
		}
		g1, dir, w, g2 := string(m[1]), string(m[2]), string(m[3]), string(m[4])
		units, err := strconv.Atoi(w)
		if err != nil {
			return nil, fmt.Errorf("Invalid units %s: %v", w, err)
		}
		if dir == "lose" {
			units = -units
		}
		if _, ok := ss[g1]; !ok {
			ss[g1] = make(map[string]int)
		}
		ss[g1][g2] = units
	}
	return ss, nil
}

func main() {
	var part1 bool
	flag.BoolVar(&part1, "part1", true, "Part 1 behavior")
	flag.Parse()
	fh := os.Stdin
	if flag.NArg() > 0 {
		if f, err := os.Open(flag.Arg(0)); err != nil {
			log.Fatalf("can't open file %s", flag.Arg(0))
		} else {
			fh = f
			defer func() { f.Close() }()
		}
	}
	ss, err := loadWeights(fh)
	if err != nil {
		log.Fatalf("%s", err)
	}
	fmt.Printf("Pairings:\n%s\n", ss.String())

}
