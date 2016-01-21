package word

import (
	//"fmt"
	"strings"
)

type Word string

func (w Word) Next() Word {
	b := []byte(w)
	for i := len(b) - 1; i >= 0; i-- {
		if b[i] == 'z' {
			b[i] = 'a'
		} else {
			b[i]++
			break
		}
	}
	return Word(b)
}

func (w Word) HasIllegalchars() bool {
	return strings.IndexAny(string(w), "iol") != -1
}

func (w Word) SequentialOf(cnt int) bool {
	ret := false
	var prev byte
	var seqcnt int
	for _, a := range []byte(w) {
		if prev == 0 {
			prev = a
		}
		//fmt.Printf("Got %d with %d %d\n", seqcnt, prev, a)
		if prev+1 == a {
			seqcnt++
			//fmt.Printf("%d %c->%c\n", seqcnt, prev, a)
			prev = a
			if seqcnt >= cnt {
				ret = true
				return ret
			}
		} else {
			seqcnt = 1
			//fmt.Printf("%d %c->%c\n", seqcnt, prev, a)
			prev = a
		}

	}
	return ret
}

func (w Word) DoubleCount() bool {
	var prev byte
	dbl := 0
	for _, a := range []byte(w) {
		if prev == 0 {
			prev = a
			continue
		}
		if prev == a {
			dbl++
			if dbl >= 2 {
				return true
			}
		}
	}
	return false
}

func (w Word) Accept() bool {
	if w.HasIllegalchars() {
		//fmt.Printf(" - failed illegal chars\n")
		return false
	}

	if !w.SequentialOf(3) {
		//fmt.Printf(" - failed sequential chars\n")
		return false
	}

	if !w.DoubleCount() {
		//fmt.Printf(" - failed double count\n")
		return false
	}

	return true
}
