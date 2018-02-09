package fibonacci

import "math/big"

// Generator returns the next permutation of the fibonacci sequence
func Generator() func() int {

	n1 := int(0)
	n2 := int(0)

	return func() (n int) {
		if n2 == 0 {
			n = 1
		} else {
			n = n1 + n2
		}
		n1 = n2
		n2 = n
		return
	}
}

// BigGenerator returns the next permutation of the fibonacci sequence
func BigGenerator() func() *big.Int {

	n1 := big.NewInt(0)
	n2 := big.NewInt(0)

	first := true

	return func() (n *big.Int) {
		if first {
			n = big.NewInt(1)
			first = false
		} else {
			n = n1.Add(n1, n2)
		}
		n1 = n2
		n2 = n
		return
	}
}
