// package main
/*

The number, 197, is called a circular prime because all rotations of the digits: 197, 971, and 719, are themselves prime.

There are thirteen such primes below 100: 2, 3, 5, 7, 11, 13, 17, 31, 37, 71, 73, 79, and 97.

How many circular primes are there below one million?

*/

package main

import (
	"fmt"
	"strconv"

	"github.com/pwmcintyre/projecteuler/functions"
	"github.com/pwmcintyre/projecteuler/prime"
)

// Run run
func Run(below int) (answer int) {

	for n := 1; n < below; n++ {
		if IsCircularPrime(n) {
			fmt.Println(n)
			answer++
		}
	}

	return
}

func IsCircularPrime(n int) bool {

	// conter to a number to get digit permutations
	nstr := strconv.Itoa(n)

	permutations := functions.Circulations(nstr)
	for p := range permutations {

		// convert to number and check if prime
		pnum, _ := strconv.Atoi(permutations[p])
		if !prime.IsPrime(pnum) {
			return false
		}

	}

	return true

}

func main() {
	fmt.Println(Run(1000000))
}
