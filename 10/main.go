package main

import (
	"bytes"
	"fmt"
	//"github.com/pkg/profile"
	"log"
	"os"
	"strconv"
)

func main() {
	reps := 40
	r := 0
	n := []byte("11")

	args := os.Args[1:]
	var err error
	if len(args) > 0 {
		reps, err = strconv.Atoi(args[0])
		if err != nil {
			log.Fatalf("%v", err)
		}
		args = args[1:]
	}
	if len(args) > 0 {
		n, args = []byte(args[0]), args[1:]
	}

	next := make([]byte, 0)
	for r < reps {
		next = []byte{}
		for len(n) > 0 {

			a := n[0]
			i := bytes.IndexFunc(n, func(r rune) bool { return a != byte(r) })
			if i == -1 {
				// will cause to break out..
				i = len(n)
				n = []byte{}
			} else {
				//fmt.Printf("n is %s, a is %c, i is %d\n", string(n), a, i)
				n = n[i:]
			}
			next = append(next, byte(i)+'0', a)
			//fmt.Printf("  next: %s\n", string(next))
		}
		r++
		n = next
		fmt.Printf("%s\n", next)
		//fmt.Printf("%d Current length is %d\n", r, len(n))
	}
}
