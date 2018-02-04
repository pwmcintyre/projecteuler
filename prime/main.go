package prime

import "math"

// prime Length returns a number of prime digits
// brute force method
func Length(length int) []int {

	primes := []int{}
	for i := 2; len(primes) < length; i++ {
		prime := true
		for j := i - 1; j > 1; j-- {
			if i%j != 0 {
				continue
			}
			prime = false
			break
		}
		if prime {
			primes = append(primes, i)
		}
	}

	return primes
}

// prime Until returns all primes <= than a given integer
// https://en.wikipedia.org/wiki/Sieve_of_Eratosthenes
func Until(until int) []int {

	if until < 2 {
		return []int{}
	}

	// golang defaults bool to false
	numbers := make([]bool, until+1)

	// √n rounded up
	end := int(math.Sqrt(float64(until)) + 0.5)
	for i := 2; i <= end; i++ {
		if numbers[i] == false {
			j := 0
			for h := 0; ; h++ {
				j = i*2 + h*i
				if j > until {
					break
				}
				numbers[j] = true
			}
		}
	}

	sequence := []int{}
	for i := 2; i < len(numbers); i++ {
		if numbers[i] == false {
			sequence = append(sequence, i)
		}
	}

	return sequence
}
