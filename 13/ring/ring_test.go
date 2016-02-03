package ring

import (
	"testing"
)

var testNormal = [][]Ring{
	{Ring{1, 2, 3}, Ring{1, 2, 3}},
	{Ring{3, 2, 1}, Ring{1, 2, 3}},
	{Ring{4, 3, 5, 7}, Ring{3, 4, 7, 5}},
}

func TestNormalize(t *testing.T) {
	for _, p := range testNormal {
		r0, normal := p[0], p[1]
		r1 := r0.Normalize()
		if len(r1) != len(normal) {
			t.Errorf("copy/normalize: %v != %v", r1, normal)
		}
		for i := range r1 {
			if r1[i] != normal[i] {
				t.Errorf("copy/normalize: %v != %v", r1, normal)
			}
		}
	}
}
