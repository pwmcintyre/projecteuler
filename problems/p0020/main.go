// package main
/*

n! means n × (n − 1) × ... × 3 × 2 × 1

For example, 10! = 10 × 9 × ... × 3 × 2 × 1 = 3628800,
and the sum of the digits in the number 10! is 3 + 6 + 2 + 8 + 8 + 0 + 0 = 27.

Find the sum of the digits in the number 100!

*/

package main

import (
	"fmt"
	"math/big"
	"strconv"

	"github.com/pwmcintyre/projecteuler/functions"
)

// Run run
func Run(n int) (answer int) {

	// get factorial in string format
	str := functions.BigFactorial(big.NewInt(int64(n))).String()

	// add the digits
	for _, e := range str {
		n, _ = strconv.Atoi(string(e))
		answer += n
	}

	return
}

func main() {
	fmt.Println(Run(100))
}
