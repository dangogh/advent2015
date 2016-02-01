package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"unicode"
)

func isNum(c rune) bool {
	return unicode.IsDigit(c)
}

func splitNumbers(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if len(data) == 0 {
		return 0, data, nil
	}

	var i int
	if rune(data[0]) == '-' || isNum(rune(data[0])) {
		i = bytes.IndexFunc(data[1:], func(b rune) bool { return !isNum(b) })
	} else {
		i = bytes.IndexFunc(data[1:], func(b rune) bool { return isNum(b) || b == '-' })
	}
	if i != -1 {
		return i + 1, data[0 : i+1], nil
	} else {
		if atEOF {
			return len(data), data, nil
		} else {
			return 0, nil, nil
		}
	}
}

func addInts(r io.Reader) int {
	s := bufio.NewScanner(r)

	s.Split(splitNumbers)
	var sum int
	for s.Scan() {
		tok := s.Bytes()
		if len(tok) == 0 {
			break
		}
		num, err := strconv.Atoi(string(tok))
		if err == nil {
			fmt.Println(num)
			sum += num
		}
	}
	return sum
}

func main() {
	f := os.Stdin
	if len(os.Args) > 1 {
		var err error
		f, err = os.Open(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
	}
	reader := bufio.NewReader(f)
	s := addInts(reader)
	fmt.Printf("%d\n", s)
}
