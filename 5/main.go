package main

import (
	"bufio"
	"fmt"
	"github.com/dangogh/advent2015/5/word"
	"log"
	"os"
)

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

func main() {
	reader := openInput()
	scanner := bufio.NewScanner(reader)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for scanner.Scan() {
		w := word.Word(scanner.Text())
		fmt.Println(w)

	}
}
