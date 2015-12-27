package word

import (
	"log"
	"testing"
)

func TestA(t *testing.T) {
	var wordTests = []struct {
		w                  string
		vowels             int
		double, restricted bool
	}{
		{"aeiou", 5, false, false},
		{"xxyyzz", 0, true, true},
	}
	for _, tt := range wordTests {
		log.Printf("%s has %d vowels, double letters? %v, restricted substrings? %v", tt.w, tt.vowels, tt.double, tt.restricted)
		w := Word(tt.w)
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

func TestB(t *testing.T) {

	var wordTests = []struct {
		w      string
		isNice bool
	}{
		{"qjhvhtzxzqqjkmpb", true},
		{"xxyxx", true},
		{"uurcxstgmygtbstg", false},
		{"ieodomkazucvgmuy", false},
	}
	for _, tt := range wordTests {
		w := Word(tt.w)
		if w.IsNice2() != tt.isNice {
			t.Errorf("%s IsNice? %v", tt.w, w.IsNice2())
		}

	}
}
