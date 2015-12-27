package word

import (
	"strings"
)

type Word string

var vowels map[rune]struct{}
var disallowed = []string{"ab", "cd", "pq", "xy"}

func init() {
	vowels = make(map[rune]struct{}, 5)
	for _, x := range "aeiou" {
		vowels[x] = struct{}{}
	}
}

func (w Word) countVowels() int {
	count := 0
	for _, c := range w {
		if _, ok := vowels[c]; ok {
			count++
		}
	}
	return count
}

func (w Word) hasDoubleLetters() bool {
	var p rune
	for _, c := range w {
		if p == c {
			return true
		}
		p = c
	}
	return false
}

func (w Word) hasDisallowedSubstring() bool {
	for _, dis := range disallowed {
		if strings.Contains(string(w), dis) {
			return true
		}
	}
	return false

}

func (w Word) IsNice() bool {
	if w.countVowels() > 2 &&
		w.hasDoubleLetters() &&
		!w.hasDisallowedSubstring() {
		return true
	}
	return false
}
