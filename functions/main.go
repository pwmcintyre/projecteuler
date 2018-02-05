package functions

import (
	"math"
	"reflect"
	"strconv"
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

func IsPalindromic(input int) bool {
	str := strconv.Itoa(input)
	for i, _ := range str {
		if str[i] != str[len(str)-1-i] {
			return false
		}
	}
	return true
}

func MaxPositiveInt(input []int) int {
	max := 0
	for _, e := range input {
		if e > max {
			max = e
		}
	}
	return max
}

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

func InArray(val interface{}, array interface{}) (exists bool, index int) {
	exists = false
	index = -1

	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)

		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) == true {
				index = i
				exists = true
				return
			}
		}
	}

	return
}
