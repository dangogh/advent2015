package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type pair struct {
	a, b string
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: ... <input>\n")
		os.Exit(1)
	}
	f, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	s := bufio.NewScanner(f)
	s.Split(bufio.ScanLines)
	sub := make([]pair, 0, 100)

	var mystr string
	for s.Scan() {
		t := s.Text()
		var a, b string
		n, err := fmt.Sscanf(t, "%s => %s", &a, &b)
		if n == 0 || err != nil {
			strings.TrimSpace(t)
			if len(t) > 0 {
				mystr = t
			}
			continue
		}
		sub = append(sub, pair{a, b})
	}

	// Now go thru the string and do the subs
	combos := map[string]struct{}{mystr: struct{}{}}

	fmt.Println("line is ", mystr)
	for _, p := range sub {
		w := strings.Replace(mystr, p.a, p.b, -1)
		fmt.Printf("Replacing %s with %s\n", p.a, p.b)
		combos[w] = struct{}{}
	}

	for s, _ := range combos {
		fmt.Println(s)
	}
}
