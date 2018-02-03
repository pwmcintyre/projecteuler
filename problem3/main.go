// Package problem3 The prime factors of 13195 are 5, 7, 13 and 29.
// What is the largest prime factor of the number 600851475143 ?

package problem3

import (
	"math"

	"github.com/pwmcintyre/projecteuler/functions"
)

// Run run
func Run(input int) int {
	factors := PrimeFactors(input)
	return functions.MaxPositiveInt(factors)
}

// PrimeFactors returns sequence of prime factors of a given integer
// https://en.wikipedia.org/wiki/Integer_factorization
// https://www.geeksforgeeks.org/print-all-prime-factors-of-a-given-number/
func PrimeFactors(n int) []int {

	save := []int{}
	for n%2 == 0 {
		save = append(save, 2)
		n /= 2
	}

	for i := 3; float64(i) <= math.Sqrt(float64(n)); i += 2 {
		for n%i == 0 {
			save = append(save, i)
			n /= i
		}
	}

	if n > 2 {
		save = append(save, n)
	}

	return save
}
