// package main
// The following iterative sequence is defined for the set of positive integers:

// n → n/2 (n is even)
// n → 3n + 1 (n is odd)

// Using the rule above and starting with 13, we generate the following sequence:

// 13 → 40 → 20 → 10 → 5 → 16 → 8 → 4 → 2 → 1
// It can be seen that this sequence (starting at 13 and finishing at 1) contains 10 terms. Although it has not been proved yet (Collatz Problem), it is thought that all starting numbers finish at 1.

// Which starting number, under one million, produces the longest chain?

// NOTE: Once the chain starts the terms are allowed to go above one million.
package main

import (
	"fmt"
)

// Run run
func Run(below int) int {

	longesti := 1
	longestlen := 0

	for i := 1; i < below; i++ {
		steps := Walk(i)
		if steps > longestlen {
			longesti = i
			longestlen = steps
		}
	}

	return longesti
}

func Walk(value int) (steps int) {
	steps = 0
	for ; value != 1; steps++ {
		if value%2 == 0 {
			value = value / 2
		} else {
			value = 3*value + 1
		}
	}
	return
}

func main() {
	fmt.Println(Run(1000000))
}
