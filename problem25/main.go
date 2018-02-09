// package main
/*

The Fibonacci sequence is defined by the recurrence relation:

Fn = Fn−1 + Fn−2, where F1 = 1 and F2 = 1.
Hence the first 12 terms will be:

F1 = 1
F2 = 1
F3 = 2
F4 = 3
F5 = 5
F6 = 8
F7 = 13
F8 = 21
F9 = 34
F10 = 55
F11 = 89
F12 = 144
The 12th term, F12, is the first term to contain three digits.

What is the index of the first term in the Fibonacci sequence to contain 1000 digits?

*/

package main

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/pwmcintyre/projecteuler/fibonacci"
)

// Run run
func Run(digits int) (answer int) {

	// create a number which has n digits
	str := "1" + strings.Repeat("0", digits-1)

	max := big.NewInt(0)
	max.SetString(str, 10)

	next := fibonacci.BigGenerator()
	n := big.NewInt(0)

	for answer = 0; n.Cmp(max) < 0; answer++ {
		n = next()
	}

	return
}

func main() {
	fmt.Println(Run(1000))
}
