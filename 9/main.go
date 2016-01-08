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
		//fmt.Printf("%s:  %s to %s = %d", line, c1, c2, dist)

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
	//fmt.Printf("cities: %v\n", cities)

	const MaxUint = ^uint(0)
	mindist := MaxUint
	mathutil.PermutationFirst(cities)

	distances := make(map[string]uint, len(cities))
	for {
		if !mathutil.PermutationNext(cities) {
			break
		}

		//fmt.Printf("Permutation %v\n", cities)
		var dist uint
		var there string
		for _, here := range cities {
			if there == "" {
				there = here
				continue
			}
			// how far from there to here
			dist += city2city[there][here]
		}
		distances[strings.Join(cities, ":")] = dist
		if dist < mindist {
			mindist = dist
		}
	}
	for c, d := range distances {
		fmt.Printf("%7d %s\n", d, c)
	}

	fmt.Printf("Min distance is %d\n", mindist)
}
