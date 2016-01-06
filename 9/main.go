package main

import (
	"bufio"
	"fmt"
	"github.com/cznic/mathutil"
	"log"
	"os"
	"strings"
)

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

		if _, ok := city2city[c1]; !ok {
			city2city[c1] = make(map[string]uint)
		}
		if _, ok := city2city[c2]; !ok {
			city2city[c2] = make(map[string]uint)
		}
		city2city[c1][c2] = dist
		city2city[c2][c1] = dist
	}
	cities := make([]string, len(city2city))
	for c, _ := range city2city {
	  cities = append(cities, c)
	}
	mathutil.PermuteFirst(cities)
	for {
	  go func() {
		cityCopy := make([]string, len(cities))
		copy(cityCopy, cities)
		var there string
		for _, here := range cityCopy {
			if there == "" {
			  there = here
			  continue
			}
			// how far from there to here
			dist += city2city[there][here]
		}
	  }
	  if !mathutil.PermuteNext(cities) {
	  	break
	  }
	}
	fmt.Printf("city map: %v\n", city2city)
}
