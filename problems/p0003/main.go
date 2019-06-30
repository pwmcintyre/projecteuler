// package main The prime factors of 13195 are 5, 7, 13 and 29.
// What is the largest prime factor of the number 600851475143 ?

package main

import (
	"github.com/pwmcintyre/projecteuler/arrays"
	"github.com/pwmcintyre/projecteuler/prime"
)

// Run run
func Run(input int) int {
	factors := prime.PrimeFactors(input)
	return arrays.MaxPositiveInt(factors)
}

func main() {

}
