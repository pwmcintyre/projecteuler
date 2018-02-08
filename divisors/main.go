package divisors

import (
	"math"
)

// Divisors returns a list of divisors of a given number
// optimised via https://stackoverflow.com/questions/171765/what-is-the-best-way-to-get-all-the-divisors-of-a-number
func Divisors(input int) []int {
	divisors := []int{}
	bigdivisors := []int{}
	limit := int(math.Sqrt(float64(input)) + 0.5)
	for i := 1; i <= limit; i++ {
		if input%i == 0 {
			divisors = append(divisors, i)
			if input != i*i {
				bigdivisors = append(bigdivisors, input/i)
			}
		}
	}
	for i := len(bigdivisors) - 1; i >= 0; i-- {
		divisors = append(divisors, bigdivisors[i])
	}
	return divisors
}

// ProperDivisors returns numbers less than n which divide evenly into n
func ProperDivisors(input int) (result []int) {
	result = Divisors(input)
	if len(result) > 0 {
		result = result[:len(result)-1]
	}
	return
}

func DivisorCount(input int) int {
	count := 0
	limit := int(math.Sqrt(float64(input)) + 1)
	for i := 1; i < limit; i++ {
		// small divisor
		if input%i == 0 {
			count++
			// big divisor
			if input != i*i {
				count++
			}
		}
	}
	return count
}
