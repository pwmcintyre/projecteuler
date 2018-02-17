// Package fibonacci is a utility to fetch permutations of the fibonacci sequence.
// More info: https://en.wikipedia.org/wiki/Fibonacci_number
// Correctness verified against https://www.bigprimes.net/archive/fibonacci/10000/
package fibonacci

import (
	"errors"
	"math/big"
)

// Fib returns a function which returns the next permutation of the fibonacci sequence.
// It will hit an int overflow after 92 permutations.
func Fib() func() (int, error) {

	n1, n2 := 0, 1

	return func() (int, error) {
		n1, n2 = n2, n1+n2
		if n1 < 0 {
			return 0, errors.New("overflow, use Bigfib")
		}
		return n1, nil
	}
}

// BigFib returns a function which returns the next permutation of the fibonacci sequence.
// Uses Big Int, so could go forever.
func BigFib() func() *big.Int {

	n1 := big.NewInt(0)
	n2 := big.NewInt(1)

	return func() *big.Int {
		n1, n2 = n2, n1.Add(n1, n2)
		return n1
	}
}
