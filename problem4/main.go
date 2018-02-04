// package main A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.
// Find the largest palindrome made from the product of two 3-digit numbers.

package main

import (
	"errors"
	"math"

	"github.com/pwmcintyre/projecteuler/functions"
)

// Run run
func Run(digits int) (int, error) {
	results, err := PalindromicNumbers(digits)
	return functions.MaxPositiveInt(results), err
}

func PalindromicNumbers(digits int) ([]int, error) {

	largest := int(math.Pow(10, float64(digits)) - 1)
	smallest := int(math.Pow(10, float64(digits-1)))

	results := []int{}

	for a := largest; a > smallest; a-- {
		for b := a - 1; b > smallest; b-- {
			result := a * b
			if functions.IsPalindromic(result) {
				results = append(results, result)
			}
		}
	}

	if len(results) == 0 {
		return nil, errors.New("no palindromic numbers found")
	}

	return results, nil
}
