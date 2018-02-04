// package main Run If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.
//
// Find the sum of all the multiples of 3 or 5 below 1000.
package main

// Input input
type Input struct {
	below, multa, multb int
}

// Run run
func Run(in Input) int {

	accumulator := 0

	for i := 0; i < in.below; i++ {
		if i%in.multa == 0 || i%in.multb == 0 {
			accumulator += i
		}
	}

	return accumulator
}
