// package main The sum of the squares of the first ten natural numbers is,

// 12 + 22 + ... + 102 = 385
// The square of the sum of the first ten natural numbers is,

// (1 + 2 + ... + 10)2 = 552 = 3025
// Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is 3025 − 385 = 2640.

// Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.

package main

// Run run
func Run(num int) int {
	a := SumOfSquares(num)
	b := SquareOfSum(num)
	return b - a
}

func SumOfSquares(numbers int) int {
	acc := 0
	for i := 1; i <= numbers; i++ {
		acc += (i * i)
	}
	return acc
}

func SquareOfSum(numbers int) int {
	acc := 0
	for i := 1; i <= numbers; i++ {
		acc += i
	}
	return acc * acc
}
