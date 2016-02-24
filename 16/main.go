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

var debug = make(chan string)

func filter(inch <-chan auntsue, key string, val int) chan auntsue {
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
				if v != val {
					// known and not same as detected
					continue
				} else {
				}
			}
			countmatch++
			outch <- s
		}
		debug <- fmt.Sprintf("%d/%d match %s=%d\n", countmatch, countall, key, val)
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
