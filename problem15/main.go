// package main
// Starting in the top left corner of a 2×2 grid, and only being able to move to the right and down, there are exactly 6 routes to the bottom right corner.
// How many such routes are there through a 20×20 grid?
package main

import (
	"fmt"
	"math/big"

	"github.com/pwmcintyre/projecteuler/factorial"
)

// Run run
// http://www.robertdickau.com/lattices.html
// (2n)! / n!^2
func Run(input uint) (answer *big.Int) {

	n := big.NewInt(int64(input))
	n2 := new(big.Int)
	n2.Set(n)
	n2 = n2.Mul(n2, big.NewInt(2))

	fact1 := factorial.BigFactorial(n2)
	fact2 := factorial.BigFactorial(n)

	fact2 = fact2.Mul(fact2, fact2)

	answer = fact1.Div(fact1, fact2)

	return
}

func main() {
	fmt.Println(Run(20))
}
