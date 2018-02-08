package numbers

import (
	"github.com/pwmcintyre/projecteuler/arrays"
	"github.com/pwmcintyre/projecteuler/divisors"
)

// PerfectNumbers returns Perfect Numbers under a given n
// https://en.wikipedia.org/wiki/Perfect_numbers
// a perfect number is a positive integer that is equal to the sum of its proper positive divisors, that is, the sum of its positive divisors excluding the number itself
func PerfectNumbers(below int) (answer []int) {

	numbers := []int{}
	for a := 2; a < below; a++ {

		if IsPerfectNumber(a) {
			numbers = append(numbers, a)
		}
	}

	return numbers
}

func IsPerfectNumber(a int) (answer bool) {

	b := arrays.Sum(divisors.ProperDivisors(a))
	c := arrays.Sum(divisors.ProperDivisors(b))

	return a != b && a == c
}
