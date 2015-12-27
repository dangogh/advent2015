package main

import (
	"bufio"
	"fmt"
	"github.com/dangogh/advent2015/5/word"
	"log"
	"os"
)

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
		fmt.Printf("Reading from %s\n", os.Args[1])
		reader = bufio.NewReader(f)
	}
	scanner := bufio.NewScanner(reader)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for scanner.Scan() {
		w := word.Word(scanner.Text())
		st := "naughty"
		if w.IsNice() {
			st = "nice"
		}
		fmt.Printf("%s is %s\n", w, st)
	}

}
