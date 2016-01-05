package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func decodeString(line string) string {
	s := line[1 : len(line)-1]
	s = strings.Replace(s, "\\\\", "\\", -1)
	s = strings.Replace(s, "\\\"", "\"", -1)
	return s
}

func encodeString(line string) string {
	s := line
	s = strings.Replace(s, "\\", "\\\\", -1)
	s = strings.Replace(s, "\"", "\\\"", -1)
	s = "\"" + s + "\""
	return s
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
		line := strings.TrimSpace(scanner.Text())
		decoded_s := decodeString(line)
		encoded_s := encodeString(line)
		log.Printf("%s: %d code chars represents %d chars and %d encoded\n", line, len(line), len(decoded_s), len(encoded_s))

		totalCode += len(line)
		totalChars += len(decoded_s)
		totalEncoded += len(encoded_s)
	}
	fmt.Printf("totalCode: %d, totalChars: %d,  solution: %d, part2: %d\n", totalCode, totalChars, totalCode-totalChars, totalEncoded-totalCode)
}
