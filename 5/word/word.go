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

/* Part 1 */
func (w Word) IsNice() bool {
	if w.countVowels() > 2 &&
		w.hasDoubleLetters() &&
		!w.hasDisallowedSubstring() {
		return true
	}
	return false
}

/* Part 2 */

/* It contains a pair of any two letters that appears at least twice in the string without
* overlapping, like xyxy (xy) or aabcdefgaa (aa), but not like aaa (aa, but it overlaps).  It
* contains at least one letter which repeats with exactly one letter between them, like xyx,
* abcdefeghi (efe), or even aaa.
 */

func (w Word) pairAppearsTwice() bool {
	s := make(map[string]struct{}, len(w))
	var prev string
	for i := 0; i < len(w)-1; i++ {
		// get next letter pair
		p := string(w[i : i+2])
		if _, ok := s[p]; ok {
			// pair appeared previously
			return true
		}
		if len(prev) != 0 {
			s[prev] = struct{}{}
		}
		prev = p
	}
	return false
}

func (w Word) repeatedLetterWithSeperation() bool {
	var c0, c1 rune
	for _, c := range w {
		if c == c1 {
			return true
		}
		c0, c1 = c, c0
	}
	return false
}

func (w Word) IsNice2() bool {
	if w.pairAppearsTwice() && w.repeatedLetterWithSeperation() {
		return true
	}
	return false
}
