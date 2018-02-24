// package main
// The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.

// Find the sum of all the primes below two million.

package main

import (
	"fmt"

	"github.com/pwmcintyre/projecteuler/prime"
)

// Run run
func Run(max int) int {

	primes := prime.Until(max)

	sum := 0
	for _, e := range primes {
		sum += e
	}

	return sum
}

func main() {
	p := Run(2000000)
	fmt.Println(p)
}
