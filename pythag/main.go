// package pythag
// https://en.wikipedia.org/wiki/Pythagorean_triple

package pythag

import (
	"math"
)

func Until(max int) [][]int {

	numbers := [][]int{}

	for a := 1; a < max; a++ {
		for b := a + 1; b < max; b++ {
			c := math.Sqrt(math.Pow(float64(a), 2.0) + math.Pow(float64(b), 2))
			if math.Remainder(c, 1) == 0 {
				c2 := int(c)
				numbers = append(numbers, []int{a, b, c2})
			}
		}
	}

	return numbers
}
