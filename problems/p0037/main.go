// package main
/*

The number 3797 has an interesting property. Being prime itself, it is possible to continuously remove digits from left to right, and remain prime at each stage: 3797, 797, 97, and 7. Similarly we can work from right to left: 3797, 379, 37, and 3.

Find the sum of the only eleven primes that are both truncatable from left to right and right to left.

NOTE: 2, 3, 5, and 7 are not considered to be truncatable primes.

*/

package main

import (
	"fmt"

	"github.com/pwmcintyre/projecteuler/functions"
	"github.com/pwmcintyre/projecteuler/prime"
)

// Run run
func Run(targetCount int) (answer int) {

	truncatablePrimes := []int{}
	for n := 10; len(truncatablePrimes) < targetCount; n++ {

		if IsTruncatablePrime(n) {
			fmt.Println(n)
			truncatablePrimes = append(truncatablePrimes, n)
		}

	}

	for _, n := range truncatablePrimes {
		answer += n
	}

	return
}

func IsTruncatablePrime(n int) bool {

	if n < 10 {
		return false
	}

	// convert to number and check if prime
	numbers := functions.TruncationsOf(n)

	for _, n := range numbers {
		if !prime.IsPrime(n) {
			return false
		}

	}

	return true

}

func main() {
	fmt.Println(Run(11))
}
