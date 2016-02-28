package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const max = 25

func volumecheck(max, curvol int, cur, rem []int, meets chan<- []int) {
	if curvol >= max {
		if curvol == max {
			meets <- cur
		}
		return
	}

	vol := curvol
	for i, v := range rem {
		a := make([]int, len(cur), len(cur)+1)
		copy(a, cur)
		a = append(a, v)
		volumecheck(max, vol+v, a, rem[i+1:], meets)
	}
}

func main() {
	if len(os.Args) < 3 {
		fmt.Printf("Usage: ... <max> <input>\n")
		os.Exit(1)
	}
	max, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
	}
	f, err := os.Open(os.Args[2])
	if err != nil {
		panic(err)
	}
	var containers []int
	s := bufio.NewScanner(f)
	s.Split(bufio.ScanLines)
	for s.Scan() {
		t := s.Text()
		n, err := strconv.Atoi(t)
		if err != nil {
			panic(err)
		}
		containers = append(containers, n)
	}
	meets := make(chan []int)
	go func() {
		defer close(meets)
		volumecheck(max, 0, []int{}, containers, meets)
	}()

	minContainers := len(containers)
	var numMin = 0
	for a := range meets {
		if len(a) < minContainers {
			numMin = 1
			minContainers = len(a)
		} else if len(a) == minContainers {
			numMin++
		}
		fmt.Printf("%+v\n", a)
	}

	fmt.Printf("%d combinations of %d containers\n", numMin, minContainers)
}
