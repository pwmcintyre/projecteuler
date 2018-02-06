// package main
// The sequence of triangle numbers is generated by adding the natural numbers. So the 7th triangle number would be 1 + 2 + 3 + 4 + 5 + 6 + 7 = 28. The first ten terms would be:

// 1, 3, 6, 10, 15, 21, 28, 36, 45, 55, ...

// Let us list the factors of the first seven triangle numbers:

//  1: 1
//  3: 1,3
//  6: 1,2,3,6
// 10: 1,2,5,10
// 15: 1,3,5,15
// 21: 1,3,7,21
// 28: 1,2,4,7,14,28
// We can see that 28 is the first triangle number to have over five divisors.

// What is the value of the first triangle number to have over five hundred divisors?
package main

import (
	"fmt"

	"github.com/pwmcintyre/projecteuler/functions"
)

// Run run
// Given a target minimum number of divisors, returns the triangle number which has at least that number of divisors, the sum value
func Run(target int) (layer int, number int) {
	for layer = 1; ; layer++ {
		number += layer
		divisors := functions.Divisors(number)
		if len(divisors) >= target {
			return
		}
	}
}

func main() {
	fmt.Println(Run(501))
}