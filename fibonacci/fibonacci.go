package fibonacci

import (
	"errors"
	"math/big"
)

// Fib returns the next permutation of the fibonacci sequence
func Fib() func() (int, error) {

	n1, n2 := 0, 1

	return func() (n int, e error) {
		n1, n2 = n2, n1+n2
		n = n1
		if n1 < 0 {
			e = errors.New("overflow, use Bigfib")
		} else {
			n = n1
		}
		return
	}
}

// BigFib returns the next permutation of the fibonacci sequence
func BigFib() func() *big.Int {

	n1 := big.NewInt(0)
	n2 := big.NewInt(1)

	return func() (n *big.Int) {
		n1, n2 = n2, n1.Add(n1, n2)
		n = n1
		return
	}
}
