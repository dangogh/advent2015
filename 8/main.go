package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func countAll(line string) (int, int) {
	s := strings.TrimSpace(line)
	countCode := 2
	countChars := 0
	// skip quotes at start and end in count
	for len(s) > 0 {
		i := strings.IndexAny(s, "\"\\")
		if i == -1 {
			countCode += len(s)
			countChars += len(s)
			break
		}

		// count chars before escaped or quote
		countCode += i
		countChars += i

		if s[i] == '"' {
			s = s[i+1:]
			continue
		}

		incr := 1
		if s[i] == '\\' {
			// look at next char
			switch s[i+1] {
			case '"':
				fallthrough
			case '\\':
				incr = 2
			case 'x':
				incr = 4
			default:
				log.Fatalf("backslash what?? %c", s[i])
			}
		}
		countCode += incr
		countChars++
		s = s[i+incr:]
	}
	return countCode, countChars
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
	scanner := bufio.NewScanner(reader)
	totalCode := 0
	totalChars := 0
	for scanner.Scan() {
		line := scanner.Text()
		codeCount, charCount := countAll(line)
		log.Printf("%s: %d code chars represents %d chars\n", line, codeCount, charCount)
		totalCode += codeCount
		totalChars += charCount
	}
	fmt.Printf("totalCode: %d, totalChars: %d,  solution: %d\n", totalCode, totalChars, totalCode-totalChars)
}
