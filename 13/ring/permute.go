package ring

import ()

// Ring is a cyclic sequence of bytes
type Ring []byte

// Permute returns channel with a sequence of unique permutations of the ring
func (r Ring) Permute() chan Ring {
	ch := make(chan Ring)
	if len(r) <= 3 {
		// only one permutation available
		go func() {
			ch <- r
			close(ch)
		}()
	} else {
		go func() {
			first := r[0]
			for p := range r[1:].Permute() {
				for i := range p {
					s := make(Ring, len(p)+1)
					copy(s, append(p[:i], append([]byte{first}, p[i:]...)...))
					ch <- s
				}
			}
			close(ch)
		}()
	}
	return ch
}
