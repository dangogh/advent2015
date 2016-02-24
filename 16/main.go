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

var detected = map[string]func(int) bool{
	"children":    func(i int) bool { return i == 3 },
	"cats":        func(i int) bool { return i > 7 },
	"samoyeds":    func(i int) bool { return i == 2 },
	"pomeranians": func(i int) bool { return i < 3 },
	"akitas":      func(i int) bool { return i == 0 },
	"vizslas":     func(i int) bool { return i == 0 },
	"goldfish":    func(i int) bool { return i < 5 },
	"trees":       func(i int) bool { return i > 3 },
	"cars":        func(i int) bool { return i == 2 },
	"perfumes":    func(i int) bool { return i == 1 },
}

var debug = make(chan string)

func filter(inch <-chan auntsue, key string, f func(int) bool) chan auntsue {
	outch := make(chan auntsue)
	go func() {
		defer close(outch)
		var countmatch, countall int
		if _, ok := detected[key]; !ok {
			fmt.Println("No key ", key)
			os.Exit(1)
		}

		for s := range inch {
			countall++
			if v, ok := s.props[key]; ok {
				if !f(v) {
					// known and not same as detected
					continue
				} else {
				}
			}
			countmatch++
			outch <- s
		}
		debug <- fmt.Sprintf("%d/%d match %s\n", countmatch, countall, key)
	}()
	return outch
}

func readInput(ch chan<- auntsue) {
	defer close(ch)
	f, _ := os.Open(os.Args[1])
	s := bufio.NewScanner(f)
	s.Split(bufio.ScanLines)
	for s.Scan() {
		t := s.Text()
		i := strings.Index(t, ":")
		id, _ := strconv.Atoi(t[4:i])
		t = t[i+1:]
		sue := auntsue{id: id, props: map[string]int{}}
		for _, pstr := range strings.Split(t, ",") {
			i := strings.Index(pstr, ":")
			k := strings.TrimSpace(pstr[:i])
			v, err := strconv.Atoi(pstr[i+2:])
			if err != nil {
				debug <- fmt.Sprintf("%v\n", err)
			}
			sue.props[k] = v
		}
		//debug <- fmt.Sprintf("readInput: %v\n", sue)
		ch <- sue
	}
}

func main() {
	go func() {
		for {
			s := <-debug
			if s == "" {
				break
			}
			fmt.Println(s)
		}
	}()

	ch := make(chan auntsue)
	go readInput(ch)

	for k, v := range detected {
		fmt.Println("Filtering ", k, v)
		ch = filter(ch, k, v)
	}
	sues := make([]auntsue, 0)
	for s := range ch {
		sues = append(sues, s)
	}

	fmt.Printf("Sue %+v\n", sues)
}
