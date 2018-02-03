// Package problem5 2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.
// What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?

package problem5

import (
	"errors"
)

// Input input
type Input struct {
	lowest, highest int
}

// Run run
func Run(input Input) (int, error) {

	for i := input.highest; i < 10000000000; i++ {
		done := true
		for x := input.lowest; x <= input.highest; x++ {
			if i%x != 0 {
				done = false
				break
			}
		}
		if done {
			return i, nil
		}
	}
	return 0, errors.New("none found")
}
