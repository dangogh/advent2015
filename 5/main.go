package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var vowels map[rune]struct{}
var disallowed []string = []string{"ab", "cd", "pq", "xy"}

func init() {
	for _, x := range "aeiou" {
		vowels[x] = struct{}{}
	}
}

func openInput() *bufio.Reader {
	var reader *bufio.Reader
	if len(os.Args) == 1 {
		reader = bufio.NewReader(os.Stdin)
	} else {
		f, err := os.Open(os.Args[1])
		defer f.Close()
		if err != nil {
			log.Fatal(err)
		}
		reader = bufio.NewReader(f)
	}
	return reader
}

func hasDoubleLetters(s string) bool {
	return false
}

func countVowels(s string) int {
	count := 0
	for _, c := range s {
		if _, ok := vowels[c]; !ok {
			count++
		}
	}
	return count
}

func hasDisallowedSubstring(s string) bool {
	for _, dis := range disallowed {
		if strings.Contains(s, dis) {
			return true
		}
	}
	return false

}

func main() {
	matcher := "000000"
	t := []string{"abcdef"}
	if len(os.Args) > 1 {
		t = os.Args[1:]
	}

	for _, prefix := range t[:] {
		i, cksum := findFirst(prefix, matcher)
		fmt.Printf("%s%d gives me %s\n", prefix, i, string(cksum))
	}
}
