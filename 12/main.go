package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func addInts(a string) int {
	s := bufio.NewScanner(strings.NewReader(a))
	split := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		fmt.Printf("string is %s\n", string(data))
		var i int
		c := data[0]
		isnum := (c == '-') || (unicode.IsDigit(rune(c)))
		for i, c = range data {
			if unicode.IsDigit(rune(c)) != isnum {
				fmt.Printf(" token %s is num? %v\n", string(data[:i]), isnum)
				return i, data[:i], nil
			}
		}
		fmt.Printf(" token %s is num? %v\n", string(data), isnum)
		return len(data), data, nil
	}
	s.Split(split)
	var sum int
	for s.Scan() {
		tok := s.Bytes()
		if len(tok) == 0 {
			//break
		}
		num, err := strconv.Atoi(string(tok))
		if err == nil {
			fmt.Println("Found number ", num)
			sum += num
		}
	}
	fmt.Printf("  last token %s\n", string(s.Bytes()))
	return sum
}

func main() {
	for _, a := range os.Args[1:] {
		s := addInts(a)
		fmt.Printf("%d\t%s\n", s, a)
	}
}
