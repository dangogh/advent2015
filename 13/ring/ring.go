package ring

import ()

// Ring describes ring of bytes
type Ring []byte

func (r Ring) minIndex() int {
	// prereq: len(r) != 0
	mini := -1
	if len(r) != 0 {
		mini = 0
	}
	for i, e := range r {
		if r[mini] > e {
			mini = i
		}
	}
	return mini
}

// Normalize creates new equivalent ring easy to compare with others
func (r Ring) Normalize() Ring {

	mini := r.minIndex()
	if mini == -1 {
		return r
	}
	c := make(Ring, len(r))
	copy(c, r)
	if mini != 0 {
		// put mini at 0
		c = append(c[mini:], c[:mini]...)
	}
	if c[1] > c[len(c)-1] {
		// reverse, but keep min at 0
		for left, right := 1, len(c)-1; left < right; left, right = left+1, right-1 {
			c[left], c[right] = c[right], c[left]
		}
	}
	return c
}
