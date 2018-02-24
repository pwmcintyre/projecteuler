// package main
// A Pythagorean triplet is a set of three natural numbers, a < b < c, for which,

// a2 + b2 = c2
// For example, 32 + 42 = 9 + 16 = 25 = 52.

// There exists exactly one Pythagorean triplet for which a + b + c = 1000.
// Find the product abc.

package main

import (
	"errors"
	"fmt"

	"github.com/pwmcintyre/projecteuler/pythag"
)

// Run run
func Run(desiredSum int) ([]int, int, error) {

	triplets := pythag.Until(desiredSum)

	for i, e := range triplets {
		sum := e[0] + e[1] + e[2]
		product := e[0] * e[1] * e[2]
		if sum == desiredSum {
			return triplets[i], product, nil
		}
	}

	return []int{}, 0, errors.New("none found")
}

func main() {
	_, p, _ := Run(1000)
	fmt.Println(p)
}
