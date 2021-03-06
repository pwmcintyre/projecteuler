package divisors

import (
	"math"
)

// GCD Greatest Common Divisor
// https://en.wikipedia.org/wiki/Greatest_common_divisor
func GCD(a int, b int) int {

	d := 0
	for a%2 == 0 && b%2 == 0 {
		a /= 2
		b /= 2
		d++
	}
	for a != b {
		if a%2 == 0 {
			a /= 2
		} else if b%2 == 0 {
			b /= 2
		} else if a > b {
			a = (a - b) / 2
		} else {
			b = (b - a) / 2
		}
	}

	result := int(float64(a) * math.Pow(2.0, float64(d)))

	return result
}
