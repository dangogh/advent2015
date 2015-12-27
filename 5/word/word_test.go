package word

import (
	"log"
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
		log.Printf("%s has %d vowels, double letters? %v, restricted substrings? %v", tt.w, tt.vowels, tt.double, tt.restricted)
		w := wword(tt.w)
		v := w.countVowels()
		d := w.hasDoubleLetters()
		r := w.hasDisallowedSubstring()

		if v != tt.vowels {
			t.Errorf("%s has %d(not %d) vowels", tt.w, v, tt.vowels)
		}
		if d != tt.double {
			t.Errorf("%s has double letters? %v", tt.w, d)
		}
		if r != tt.restricted {
			t.Errorf("%s has restricted substrings? %v", tt.w, r)
		}
	}
}
