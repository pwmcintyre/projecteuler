// package main
/*

Surprisingly there are only three numbers that can be written as the sum of fourth powers of their digits:

1634 = 1^4 + 6^4 + 3^4 + 4^4
8208 = 8^4 + 2^4 + 0^4 + 8^4
9474 = 9^4 + 4^4 + 7^4 + 4^4
As 1 = 1^4 is not a sum it is not included.

The sum of these numbers is 1634 + 8208 + 9474 = 19316.

Find the sum of all the numbers that can be written as the sum of fifth powers of their digits.

*/

package main

import (
	"fmt"
)

// Run run
func Run(below, power int) (answer int) {

	for _, n := range SumOfPower(below, power) {
		answer += n
	}

	return
}

// SumOfPower returns a list of integers where IsSumOfPower(N) == true and N < below
func SumOfPower(below, power int) (numbers []int) {

	// skip 1
	for i := 2; i < below; i++ {
		if IsSumOfPower(i, power) {
			numbers = append(numbers, i)
		}
	}

	return

}

// IsSumOfPower returns true if the power of each digit equals the whole
// example. 9474 = 9^4 + 4^4 + 7^4 + 4^4 == true
func IsSumOfPower(input int, power int) bool {

	number := input

	// iterate over each digit,
	// trimming off the least sig. integer each iteration
	sum := 0
	for number > 0 {

		n := number % 10
		number /= 10

		// ghetto power operator
		p := n
		for i := 1; i < power; i++ {
			p *= n
		}

		sum += p

	}

	return sum == input

}

func main() {
	fmt.Println(Run(1000000, 5))
}
