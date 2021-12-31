package lcg

// IsPrime implements Miller-Rabin from https://en.wikipedia.org/wiki/Miller%E2%80%93Rabin_primality_test
func IsPrime(n int) bool {
	if n == 2 || n == 3 {
		return true
	}
	if n%2 == 0 {
		return false
	}

	// factor our powers of 2 from n-1 to construct n = 2^r*d+1
	r := 0
	d := n - 1
	for d%2 == 0 {
		r++
		d /= 2
	}

	// witness loop, these factors should be good up to > 2^64
witness:
	for _, a := range []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41} {
		if a > n-2 {
			continue
		}
		x := 1
		for i := 0; i < d; i++ {
			x *= a
			x %= n
		}
		if x == 1 || x == n-1 {
			continue witness
		}
		for i := 0; i < r; i++ {
			x = (x * x) % n
			if x == n-1 {
				continue witness
			}
		}
		return false
	}
	return true
}

// PrimeFactors returns the non-duplicate prime factors for a number, with a special case of returning [1] when n=1
func PrimeFactors(n int) []int {
	var factors []int
	for i := 2; i <= n/2; i++ {
		if IsPrime(i) && n%i == 0 {
			factors = append(factors, i)
		}
	}
	if n == 1 || IsPrime(n) {
		factors = append(factors, n)
	}
	return factors
}
