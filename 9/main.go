package main

import (
	"bufio"
	"fmt"
	"github.com/cznic/mathutil"
	"log"
	"os"
	"strings"
)

type cityList []string

func (c cityList) Len() int {
	return len(c)
}

func (c cityList) Less(i, j int) bool {
	sl := []string(c)
	return sl[i] < sl[j]
}

func (c cityList) Swap(i, j int) {
	sl := []string(c)
	sl[i], sl[j] = sl[j], sl[i]
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
		//fmt.Printf("Reading from %s\n", infile)
		reader = bufio.NewReader(f)
	}
	city2city := make(map[string]map[string]uint)

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		var c1, c2 string
		var dist uint
		n, err := fmt.Sscanf(line, "%s to %s = %d", &c1, &c2, &dist)
		if n != 3 {
			log.Fatalf("%d items read from %s: %v", n, line, err)
		}
		fmt.Printf("%s:  %s to %s = %d", line, c1, c2, dist)

		if _, ok := city2city[c1]; !ok {
			city2city[c1] = make(map[string]uint)
		}
		if _, ok := city2city[c2]; !ok {
			city2city[c2] = make(map[string]uint)
		}
		city2city[c1][c2] = dist
		city2city[c2][c1] = dist
	}

	cities := make(cityList, 0, len(city2city))
	for c := range city2city {
		cities = append(cities, c)
	}
	fmt.Printf("cities: %v\n", cities)

	// channel for permutations
	permch := make(chan cityList)
	mathutil.PermutationFirst(cities)

	go func() {
		cc := cities
		for {
			c := make(cityList, len(cc))
			copy(c, cc)
			permch <- c
			if !mathutil.PermutationNext(c) {
				break
			}
			cc = c
		}
	}()

	const MaxUint = ^uint(0)
	mindist := MaxUint

	ch := make(chan uint)
	for p := range permch {
		fmt.Printf("Permutation %v\n", p)
		go func(p cityList) {
			var dist uint
			var there string
			for _, here := range p {
				if there == "" {
					there = here
					continue
				}
				// how far from there to here
				dist += city2city[there][here]
				if dist > mindist {
					// give up without sending back value
					return
				}
			}
			// completed -- return distance calculated
			ch <- dist
		}(p)
	}

	for dist := range ch {
		if dist < mindist {
			mindist = dist
		}
	}

	fmt.Printf("Min distance is %d\n", mindist)
}
