package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"

	"./ring"
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

		//fmt.Println(t)
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

func getHappiness(r []byte, weights map[string]map[string]int, idmap []string) int {

	var sum int
	for i, id2 := range r {
		id1 := r[0]
		if i < len(r)-1 {
			id1 = r[i+1]
		}
		name1, name2 := idmap[id1], idmap[id2]
		//fmt.Printf("%7s->%7s:\t%d\n%7s->%7s\t%d\n", name1, name2, ws[name1][name2], name2, name1, ws[name2][name1])
		sum += weights[name1][name2] + weights[name2][name1]
	}
	return sum
}

func report(maxh int, arr []byte, idmap []string) {
	var a []string
	for _, i := range arr {
		a = append(a, idmap[i])
	}
	fmt.Println(maxh, " ", strings.Join(a, "<->"))
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
	ws, err := loadWeights(fh)
	if err != nil {
		log.Fatalf("%s", err)
	}
	fmt.Printf("Pairings:\n%s\n", ws.String())

	r := make(ring.Ring, 0, len(ws))
	idmap := make([]string, 0, len(ws))

	var i byte
	for name := range ws {
		idmap = append(idmap, name)
		r = append(r, i)
		i++
	}
	maxhappiness := math.MinInt64
	var happiestarrangement []byte
	for r0 := range r.Permute() {
		happiness := getHappiness(r0, ws, idmap)
		//fmt.Println(r0)
		if happiness > maxhappiness {
			maxhappiness = happiness
			happiestarrangement = r0
		}
	}
	report(maxhappiness, happiestarrangement, idmap)

	// Now insert myself
	idmap = append(idmap, "Dan")
	r0 := happiestarrangement
	first := byte(len(r0))
	maxh := math.MinInt64
	ws["Dan"] = make(map[string]int)
	var happarr []byte
	for i := range r0 {
		s := make(ring.Ring, len(r0)+1)
		copy(s, append(r0[:i], append([]byte{first}, r0[i:]...)...))
		h := getHappiness(s, ws, idmap)
		if h > maxh {
			maxh = h
			happarr = s
		}
	}
	fmt.Printf("Pairings:\n%s\n", ws.String())
	report(maxh, happarr, idmap)
}
