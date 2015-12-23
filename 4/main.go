package main

import (
	"bufio"
	"crypto/md5"
	"fmt"
	_ "io"
	"log"
	"os"
	"strconv"
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

func findFirst(prefix, matcher string) (int, string) {
	for i := 0; ; /*forever*/ i++ {
		ii := prefix + strconv.Itoa(i)
		cksum_b := md5.Sum([]byte(ii))
		cksum := fmt.Sprintf("%x", cksum_b[:])

		if matcher == string(cksum[:len(matcher)]) {
			return i, cksum
		}
	}
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
