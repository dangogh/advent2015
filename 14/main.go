package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
)

type reindeer struct {
	Name     string
	Speed    int
	Duration int
	Rest     int
	Points   int
}

func (r reindeer) DistanceAt(time int) int {
	i := time / (r.Duration + r.Rest)
	rem := time % (r.Duration + r.Rest)
	fmt.Printf("at %d,  %d intervals with %d remaining\n", time, i, rem)
	d := i * r.Duration * r.Speed
	if rem > r.Duration {
		rem = r.Duration
	}
	d += rem * r.Speed
	return d
}

func loadTeam(r io.Reader) ([]*reindeer, error) {
	re := regexp.MustCompile(`(\w+) can fly (\d+) km/s for (\d+) seconds, but then must rest for (\d+) seconds.`)
	var team []*reindeer
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanLines)
	for s.Scan() {
		t := s.Text()

		//fmt.Println(t)
		m := re.FindSubmatch([]byte(t))
		if m == nil {
			return nil, fmt.Errorf("no match on line: %s", t)
		}
		speed, err := strconv.Atoi(string(m[2]))
		if err != nil {
			return nil, fmt.Errorf("speed not a number? %s", m[2])
		}
		duration, err := strconv.Atoi(string(m[3]))
		if err != nil {
			return nil, fmt.Errorf("duration not a number? %s", m[3])
		}
		rest, err := strconv.Atoi(string(m[4]))
		if err != nil {
			return nil, fmt.Errorf("rest not a number? %s", m[4])
		}

		r := reindeer{Name: string(m[1]), Speed: speed, Duration: duration, Rest: rest}
		team = append(team, &r)
	}
	return team, nil
}

type trial struct {
	R        reindeer
	Distance int
}

func travel(team []*reindeer, totalTime int) chan trial {
	ch := make(chan trial)
	for _, r := range team {
		rr := r
		go func() {
			time := totalTime
			var dist int
			for time > 0 {
				if time < rr.Duration {
					dist += rr.Speed * time
					break
				}
				dist += rr.Speed * rr.Duration
				time -= (rr.Duration + rr.Rest)
			}
			ch <- trial{*rr, dist}
		}()
	}
	return ch
}

func main() {
	fh := os.Stdin
	flag.Parse()
	if flag.NArg() > 0 {
		if f, err := os.Open(flag.Arg(0)); err != nil {
			log.Fatalf("can't open file %s", flag.Arg(0))
		} else {
			fh = f
			defer func() { f.Close() }()
		}
	}

	team, err := loadTeam(fh)
	if err != nil {
		log.Fatalf("%s", err)
	}
	fmt.Printf("Team: %#v\n", team)
	ch := travel(team, 2503)
	max := 0
	for i := 0; i < len(team); i++ {
		r := <-ch
		fmt.Printf("%s traveled %d km\n", r.R.Name, r.Distance)
		if max < r.Distance {
			max = r.Distance
		}
	}
	fmt.Println("Max distance was ", max)

	for i := 1; i <= 2503; i++ {
		max := 0
		var lead *reindeer
		for _, r := range team {
			d := r.DistanceAt(i)
			if d > max {
				max = d
				lead = r
			}
		}
		fmt.Printf("%s has the lead at %d (%d)\n", lead.Name, i, max)
		(*lead).Points++
	}
	for _, r := range team {
		fmt.Printf("%+v\n", r)
	}
}
