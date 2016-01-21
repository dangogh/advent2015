package word

import (
	"testing"
)

var testMap = map[string][]bool{
	"i":        {false, false, false},
	"a":        {false, true, false},
	"hijklmmn": {true, false, false},
	"abbceffg": {false, true, true},
	"abbcegjk": {false, true, false},
	"abcdffaa": {true, true, true},
	"ghkaabcc": {true, true, true},
}

func TestSequentialOf(t *testing.T) {
	for k, res := range testMap {
		w := Word(k)
		if res[0] != w.SequentialOf(3) {
			t.Errorf("%s has sequential characters: %v", k, w.SequentialOf(3))
		}
	}
}

func TestIllegalchars(t *testing.T) {

	for k, res := range testMap {
		w := Word(k)
		if res[1] != !w.HasIllegalchars() {
			t.Errorf("%s has illegal characters: %v", k, w.HasIllegalchars())
		}
	}
}
