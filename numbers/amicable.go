package numbers

import (
	"github.com/pwmcintyre/projecteuler/arrays"
	"github.com/pwmcintyre/projecteuler/divisors"
)

// AmicableNumbers returns Amicable Numbers under a given n
// https://en.wikipedia.org/wiki/Amicable_numbers
func AmicableNumbers(below int) (answer []int) {

	numbers := []int{}
	for a := 2; a < below; a++ {

		if IsAmicableNumber(a) {
			numbers = append(numbers, a)
		}
	}

	return numbers
}

func IsAmicableNumber(a int) (answer bool) {

	b := arrays.Sum(divisors.ProperDivisors(a))
	c := arrays.Sum(divisors.ProperDivisors(b))

	return a != b && a == c
}
