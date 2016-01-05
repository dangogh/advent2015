package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func decodeString(line string) string {
	s := line
	s = strings.Replace(s, "\\\\", "\\", -1)
	s = strings.Replace(s, "\\\"", "\"", -1)
	return s
}

func encodeString(line string) string {
	s := line
	s = strings.Replace(s, "\\", "\\\\", -1)
	s = strings.Replace(s, "\"", "\\\"", -1)
	return s
}

func countAll(line string) (int, int, int) {
	countCode := 2
	countChars := 0
	countEncoded := countCode*2 + 2
	// skip quotes at start and end in count
	s := line
	for len(s) > 0 {
		i := strings.IndexAny(s, "\"\\")
		if i == -1 {
			countCode += len(s)
			countChars += len(s)
			countEncoded += len(s)
			break
		}

		// count chars before escaped or quote
		countCode += i
		countChars += i
		countEncoded += i

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
		countCode += incr * 2
		s = s[i+incr:]
	}
	return countCode, countChars, countEncoded
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
	totalEncoded := 0
	for scanner.Scan() {
		line := scanner.Text()
		line := strings.TrimSpace(line)
		decoded_s := decoded(line)
		encoded_s := enccoded(line)
		log.Printf("%s: %d code chars represents %d chars and %d encoded\n", line, len(line), len(decoded_s), len(encoded_s))

		codeCount, charCount, countEncoded := countAll(line)
		log.Printf("%s: %d code chars represents %d chars and %d encoded\n", line, codeCount, charCount, countEncoded)
		totalCode += codeCount
		totalChars += charCount
		totalEncoded += countEncoded
	}
	fmt.Printf("totalCode: %d, totalChars: %d,  solution: %d, part2: %d\n", totalCode, totalChars, totalCode-totalChars, totalEncoded-totalCode)
}
