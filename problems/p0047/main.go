// package main
/*

The first two consecutive numbers to have two distinct prime factors are:

14 = 2 × 7
15 = 3 × 5

The first three consecutive numbers to have three distinct prime factors are:

644 = 2² × 7 × 23
645 = 3 × 5 × 43
646 = 2 × 17 × 19.

Find the first four consecutive integers to have four distinct prime factors each. What is the first of these numbers?

*/

package main

import (
	"fmt"

	"github.com/pwmcintyre/projecteuler/arrays"
	"github.com/pwmcintyre/projecteuler/prime"
)

// Run run
func Run(target int) (answer int) {

	consecutives := []int{}
	for n := 10; true; n++ {

		// prime factors
		pfactors := prime.PrimeFactors(n)

		// distinct factors
		factors := arrays.UniqueInt(pfactors)

		// must have n factors
		if len(factors) != target {

			// reset and continue
			consecutives = []int{}
			continue

		}

		consecutives = append(consecutives, n)

		// must be consecutive
		if len(consecutives) == target {
			break
		}

	}

	fmt.Println(consecutives)
	return consecutives[0]
}

func main() {
	fmt.Println(Run(987654321))
}
