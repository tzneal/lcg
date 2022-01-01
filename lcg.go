package lcg

import (
	"math/rand"
)

// LCG implements a linear congruential generator from https://en.wikipedia.org/wiki/Linear_congruential_generator
// This is not intended to generate data that will pass statistical tests, but to instead provide a pseudo-random
// way of iterating over a large sequence.
//
// Suppose you want to iterate from 1 to 1-million in a non-sequential order.  You could generate those numbers,
// shuffle them somehow, etc. but you now have to store them.  An LCG with period of 1-million provides a similar
// method of iteration.
type LCG struct {
	x int // current value

	x0 int // seed
	a  int // multiplier
	m  int // modulus
	c  int // increment
}

// NewLCG constructs a new LCG with the range [0,m) using a fixed seed random source.
func NewLCG(m int) (*LCG, error) {
	rs := rand.NewSource(42)
	return NewLCGWithSource(m, rs)
}

// NewLCG constructs a new LCG with the range [0,m) using a given random source for constructing the LCG.
func NewLCGWithSource(m int, rs rand.Source) (*LCG, error) {
	rng := rand.New(rs)
	if m <= 2 {
		return &LCG{
			m: m,
			x: -1,
			c: 1,
			a: 1,
		}, nil
	}
	lcg := &LCG{
		m: m,
		x: -1,
	}

	// When c != 0, correctly chosen parameters allow a period equal to m, for all seed
	// values. This will occur if and only if:
	//
	//  m and c are relatively prime,
	//  a − 1 is divisible by all prime factors of m
	//  a − 1 is divisible by 4 if m is divisible by 4
	tries := 10
	for {
		// try to find a c that is relatively prime with m
		c := rng.Intn(m-2) + 1

		// This isn't part of the algorithm, but the sequences seem to be 'more random looking' if we keep c towards
		// the middle of m.  We only try a few times though so we can at least generate some sequence
		if (c > m-5 || c <= 5) && tries > 0 {
			tries--
			continue
		}
		if GCD(c, m) != 1 {
			continue
		}

		var a int
		mpf := PrimeFactors(m)
	inner:
		for {
			a = rng.Intn(m-1) + 1
			if a == 1 && tries > 0 {
				tries--
				continue inner
			}

			for _, p := range mpf {
				if (a-1)%p != 0 {
					continue inner
				}
			}
			if m%4 == 0 && (a-1)%4 != 0 {
				continue inner
			}
			break
		}
		lcg.c = c
		lcg.a = a
		break
	}
	return lcg, nil
}

// Done returns true if the LCG has cycled from the initial state.  Calling Next() will restart the next cycle.
func (l *LCG) Done() bool {
	return l.x == l.x0
}

// Next returns the next number in the LCG sequence.
func (l *LCG) Next() int {
	if l.x == -1 {
		l.x = 0
	}
	l.x = (l.a*l.x + l.c) % l.m
	return l.x
}
