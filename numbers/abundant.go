package numbers

import (
	"github.com/pwmcintyre/projecteuler/arrays"
	"github.com/pwmcintyre/projecteuler/divisors"
)

// AbundantNumbers returns Abundant Numbers under a given n
// https://en.wikipedia.org/wiki/Abundant_number
// In number theory, an abundant number or excessive number is a number for which the sum of its proper divisors is greater than the number itself. The integer 12 is the first abundant number. Its proper divisors are 1, 2, 3, 4 and 6 for a total of 16. The amount by which the sum exceeds the number is the abundance. The number 12 has an abundance of 4, for example.
func AbundantNumbers(below int) (answer []int) {

	numbers := []int{}
	for a := 2; a < below; a++ {

		if IsAbundantNumber(a) {
			numbers = append(numbers, a)
		}
	}

	return numbers
}

func IsAbundantNumber(a int) (answer bool) {

	b := arrays.Sum(divisors.ProperDivisors(a))

	return b > a
}

// could also produce Sum-of if needed
func NotSumOfAbundantNumbers(below int) (answer []int) {

	abundant := AbundantNumbers(below)

	sumOfAbundant := [][]int{}

	// golang defaults bool to false
	isSumOfAbundant := make([]bool, below)

	count := 0
	for n := 1; n < below; n++ {
	Loop:
		for i, a := range abundant {
			for j := i; j < len(abundant); j++ {
				b := abundant[j]
				sum := a + b
				if sum > n {
					break
				} else if a+b == n {
					sumOfAbundant = append(sumOfAbundant, []int{a, b})
					isSumOfAbundant[n] = true
					count++
					break Loop
				}
			}
		}
	}

	answer = make([]int, below-count)
	i := 0
	for n, j := range isSumOfAbundant {
		if !j {
			answer[i] = n
			i++
		}
	}

	return
}
