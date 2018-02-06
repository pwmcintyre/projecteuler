package factorial

import (
	"math/big"
)

func Factorial(n uint) (result uint) {
	if n > 0 {
		result = n * Factorial(n-1)
		return result
	}
	return 1
}

func BigFactorial(n *big.Int) (result *big.Int) {

	b := big.NewInt(0)
	c := big.NewInt(1)

	// n < 0
	if n.Cmp(b) == -1 {
		result = big.NewInt(1)
	}

	// if n == 0
	if n.Cmp(b) == 0 {
		result = big.NewInt(1)
	} else { // if n > 0

		result = new(big.Int)
		result.Set(n)

		n.Sub(n, c)           // n - 1
		n = BigFactorial(n)   // (n-1)!
		result.Mul(result, n) // n * (n - 1)!
	}
	return
}
