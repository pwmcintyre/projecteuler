// package main
/*

Starting with the number 1 and moving to the right in a clockwise direction a 5 by 5 spiral is formed as follows:

21 22 23 24 25
20  7  8  9 10
19  6  1  2 11
18  5  4  3 12
17 16 15 14 13

It can be verified that the sum of the numbers on the diagonals is 101.

What is the sum of the numbers on the diagonals in a 1001 by 1001 spiral formed in the same way?

*/

package main

import (
	"fmt"
)

// Run run
func Run(length int) (answer int) {

	max := length * length
	answer = 1

	for diameter := 1; diameter <= length; diameter += 2 {

	}

	return
}

func main() {
	fmt.Println(Run(1001, 1001))
}
