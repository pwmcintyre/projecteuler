// package main
/*

We shall say that an n-digit number is pandigital if it makes use of all the digits 1 to n exactly once. For example, 2143 is a 4-digit pandigital and is also prime.

What is the largest n-digit pandigital prime that exists?

*/

package main

import (
	"fmt"

	"github.com/pwmcintyre/projecteuler/prime"
)

// Run run
func Run(below int) (answer int) {

	// warm the cache
	fmt.Println("fetching primes")
	primes := prime.SieveOfEratosthenes(below)

	for n := len(primes) - 1; n >= 0; n-- {

		if n%100000 == 0 {
			fmt.Println(primes[n])
		}

		if IsPandigital(primes[n]) {
			return primes[n]
		}

	}

	return
}

func IsPandigital(n int) bool {

	keys := make(map[int]bool)
	list := []int{}

	// build a map and array of digits
	var ii int
	for ii = n; ii > 0; ii = ii / 10 {

		digit := ii % 10

		keys[digit] = true
		list = append(list, digit)

	}

	// ensure each n digits are used
	for i := len(list); i > 0; i-- {
		if !keys[i] {
			return false
		}
	}

	return true

}

func main() {
	fmt.Println(Run(987654321))
}
