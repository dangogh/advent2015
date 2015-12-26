package word

import (
	"testing"
)

var wordTests = []struct {
	w                  string
	vowels             int
	double, restricted bool
}{
	{"aeiou", 5, false, false},
	{"xxyyzz", 0, true, true},
}

func TestA(t *testing.T) {
	for _, tt := range wordTests {
		if wword(tt.w).countVowels() != tt.vowels {
			t.Errorf("%s has %d vowels", tt.w, tt.vowels)
		}
		if wword(tt.w).hasDoubleLetters() != tt.double {
			d := ""
			if !tt.double {
				d = "no "
			}
			t.Errorf("%s has %sdouble letters", tt.w, d)
		}
		if wword(tt.w).hasDisallowedSubstring() != tt.restricted {
			t.Errorf("%s has restricted substrings", tt.w)
		}
	}
}
