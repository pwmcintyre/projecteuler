// package main
/*

A permutation is an ordered arrangement of objects. For example, 3124 is one possible permutation of the digits 1, 2, 3 and 4. If all of the permutations are listed numerically or alphabetically, we call it lexicographic order. The lexicographic permutations of 0, 1 and 2 are:

012   021   102   120   201   210

What is the millionth lexicographic permutation of the digits 0, 1, 2, 3, 4, 5, 6, 7, 8 and 9?

*/

package main

import (
	"fmt"
	"math"
	"strconv"

	"github.com/pwmcintyre/projecteuler/functions"
)

// Run run
func Run(below, index int) (answer string) {

	digits := GenDigits(below)
	answer = NthPermutationOfOrderedArrangement(digits, index)

	return
}

// had a 1am epiphany and wrote this on my phone

/*

N current digit
M digits
D desired permutation

For n, 0 to M
Make a slice without previous numbers, new m
Digit[ ( floor ( D / (m-n-1)! ) + N ) % m ]

*/

// .... it works
func NthPermutationOfOrderedArrangement(digits []int, index int) (answer string) {

	length := len(digits)
	for i := 0; i < length-1; i++ {

		a := len(digits) - 1
		afact := int(functions.Factorial(uint(a)))

		j := int(math.Floor(float64(index) / float64(afact)))

		idx := j % len(digits)
		digit := digits[idx]

		answer = answer + strconv.Itoa(digit)

		// remove this digit
		digits = append(digits[:idx], digits[idx+1:]...)
	}

	answer = answer + strconv.Itoa(digits[0])

	return
}

func GenDigits(below int) (digits []int) {
	digits = make([]int, below)
	for i := range digits {
		digits[i] = i
	}
	return
}

func main() {
	fmt.Println(Run(10, 1000000))
}
