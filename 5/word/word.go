package word

import (
	"strings"
)

type wword string

var vowels map[rune]struct{}
var disallowed = []string{"ab", "cd", "pq", "xy"}

func init() {
	vowels = make(map[rune]struct{}, 5)
	for _, x := range "aeiou" {
		vowels[x] = struct{}{}
	}
}

func (w wword) countVowels() int {
	count := 0
	for _, c := range w {
		if _, ok := vowels[c]; !ok {
			count++
		}
	}
	return count
}

func (w wword) hasDoubleLetters() bool {
	return false
}

func (w wword) hasDisallowedSubstring() bool {
	for _, dis := range disallowed {
		if strings.Contains(string(w), dis) {
			return true
		}
	}
	return false

}
