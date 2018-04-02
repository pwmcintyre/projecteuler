// package main
/*

145 is a curious number, as 1! + 4! + 5! = 1 + 24 + 120 = 145.

Find the sum of all numbers which are equal to the sum of the factorial of their digits.

Note: as 1! = 1 and 2! = 2 are not sums they are not included.

*/

package main

import (
	"fmt"

	"github.com/pwmcintyre/projecteuler/functions"
)

// Run run
func Run(below int) (answer int) {

	for _, n := range SumOfFactorials(below) {
		answer += n
	}

	return
}

// SumOfFactorials returns a list of integers where IsSumOfFactorials(N) == true and N < below
func SumOfFactorials(below int) (numbers []int) {

	// skip 1,2
	for i := 3; i < below; i++ {
		if IsSumOfFactorials(i) {
			numbers = append(numbers, i)
		}
	}

	return

}

// IsSumOfFactorials returns true if the factorial of each digit equals the whole
// example. 145 = 1! + 4! + 5! = 1 + 24 + 120 = 145 == true
func IsSumOfFactorials(input int) bool {

	number := input

	// iterate over each digit,
	// trimming off the least sig. integer each iteration
	sum := uint(0)
	for number > 0 {

		n := uint(number % 10)
		number /= 10

		sum += functions.Factorial(n)

	}

	return sum == uint(input)

}

func main() {
	// fmt.Println(Run(1<<63 - 1)) // too big :(
	fmt.Println(Run(100000))
}
