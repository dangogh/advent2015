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

var seen = make(map[string]struct{})

func generate(word string, subs []pair) []string {
	res := make([]string, 0, len(subs))
	for _, p := range subs {
		j := strings.Index(word, p.a)
		if j == -1 {
			// no sub -- skip
			continue
		}
		var w string
		w = word[:j] + p.b + word[j+len(p.a):]
		if _, ok := seen[w]; ok {
			// already seen -- skip
			continue
		}
		seen[w] = struct{}{}
		res = append(res, w)
	}
	return res
}

func solveFormula(curlevel int, curset []string, target string, subs []pair) int {
	fmt.Println("Level is ", curlevel)
	var nextset []string
	for _, cur := range curset {
		if cur == target {
			return curlevel
		}
		nextset = append(nextset, generate(cur, subs)...)
	}
	for _, w := range nextset {
		if w == target {
			return curlevel
		}
	}
	return solveFormula(curlevel+1, nextset, target, subs)
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

	level := solveFormula(0, []string{"e"}, formula, subs)
	fmt.Println(level, " transformations to make target molecule")
}
