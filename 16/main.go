package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type auntsue struct {
	id    int
	props map[string]int
}

var detected = map[string]int{
	"children":    3,
	"cats":        7,
	"samoyeds":    2,
	"pomeranians": 3,
	"akitas":      0,
	"vizslas":     0,
	"goldfish":    5,
	"trees":       3,
	"cars":        2,
	"perfumes":    1,
}

func main() {
	ch0 := make(chan auntsue)
	debug := make(chan string)
	go func() {
		fmt.Println(<-debug)
	}()

	go func() {
		f, _ := os.Open("input.txt")
		s := bufio.NewScanner(f)
		s.Split(bufio.ScanLines)
		for s.Scan() {
			t := s.Text()
			i := strings.Index(t, ":")
			id, _ := strconv.Atoi(t[4:i])
			t = t[i+1:]
			s := auntsue{id: id, props: make(map[string]int)}
			for _, pstr := range strings.Split(t, ",") {
				i := strings.Index(pstr, ":")
				k := pstr[:i]
				v, err := strconv.Atoi(pstr[i+2:])
				if err != nil {
					debug <- fmt.Sprintf("%v\n", err)
				}
				s.props[k] = v
			}
			ch0 <- s
		}
		close(ch0)
	}()

	for k, v := range detected {
		ch1 := make(chan auntsue)
		go func() {
			for s := range ch0 {
				debug <- fmt.Sprintf("%s=%d for %+v?\n", k, v, s)
				if n, ok := s.props[k]; ok {
					if n == v {
						ch1 <- s
					}
				}
			}
			close(ch1)
		}()
		ch0 = ch1
	}
	for s := range ch0 {
		fmt.Printf("Sue %+v\n", s)
	}
}
