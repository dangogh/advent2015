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
	combos := make(map[string]struct{})

	fmt.Println("line is ", mystr)
	for i, p := range sub {
		fmt.Printf("%d %+v\n", i, p)
		var cur int
		for cur < len(mystr) {
			j := strings.Index(mystr[cur:], p.a)
			if j == -1 {
				break
			}
			cur += j
			w := mystr[:cur] + p.b + mystr[cur+len(p.a):]
			fmt.Printf("Replacing %s with %s\n", p.a, p.b)
			fmt.Println(" got ", w)
			combos[w] = struct{}{}
			cur += len(p.a)
			fmt.Println(i, " cur is ", cur, " remaining ", mystr[cur:])
		}
	}

	for s := range combos {
		fmt.Println(s)
	}
	fmt.Println(len(combos), " molecules")
}
