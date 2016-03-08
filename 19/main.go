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

func readFormula(fn string) ([]pair, string, error) {
	f, err := os.Open(fn)
	if err != nil {
		return nil, "", err
	}
	s := bufio.NewScanner(f)
	s.Split(bufio.ScanLines)
	subs := make([]pair, 0, 100)

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
		subs = append(subs, pair{a, b})
	}
	return subs, mystr, nil
}

func uniqueCombinations(subs []pair, formula string) map[string]struct{} {
	combos := make(map[string]struct{})

	fmt.Println("line is ", formula)
	for i, p := range subs {
		fmt.Printf("%d %+v\n", i, p)
		var cur int
		for cur < len(formula) {
			j := strings.Index(formula[cur:], p.a)
			if j == -1 {
				break
			}
			cur += j
			w := formula[:cur] + p.b + formula[cur+len(p.a):]
			fmt.Printf("Replacing %s with %s\n", p.a, p.b)
			fmt.Println(" got ", w)
			combos[w] = struct{}{}
			cur += len(p.a)
			fmt.Println(i, " cur is ", cur, " remaining ", formula[cur:])
		}
	}
	return combos
}

func solveFormula(level int, start, target string, subs []pair) int {
	if target == start {
		return level
	}
	var seen = make(map[string]struct{})
	for _, p := range subs {
		i := strings.Index(target, p.b)
		if i == -1 {
			// not found -- skip
			continue
		}
		w := target[:i] + p.a + target[i+len(p.b):]
		if _, ok := seen[w]; ok {
			continue
		}
		seen[w] = struct{}{}
		foundlevel := solveFormula(level+1, start, w, subs)
		if foundlevel != -1 {
			return foundlevel
		}
	}
	return -1
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: ... <input>\n")
		os.Exit(1)
	}
	subs, formula, err := readFormula(os.Args[1])

	if err != nil {
		panic(err)
	}

	if false {
		combos := uniqueCombinations(subs, formula)
		for s := range combos {
			fmt.Println(s)
		}
		fmt.Println(len(combos), " molecules")
	}

	level := solveFormula(0, "e", formula, subs)
	fmt.Println(level, " transformations to make target molecule")
}
