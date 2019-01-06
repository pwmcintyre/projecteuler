package prime

import (
	"math"
)

// cached primes
var primes = []int{2, 3, 5, 7, 11, 13}

func purgeCache() {
	primes = []int{}
}

// prime Length returns a number of prime digits
// brute force method
func Length(length int) []int {

	for i := primes[len(primes)-1] + 1; len(primes) < length; i++ {
		if IsPrime(i) {
			primes = append(primes, i)
		}
	}

	return primes[0:length]
}

// prime Until returns all primes <= than a given integer
func Until(until int) []int {

	if until < 2 {
		return []int{}
	}

	// use cache if we have enough already
	if primes[len(primes)-1] > until {
		for i, e := range primes {
			if e > until {
				return primes[0:i]
			}
		}
	}

	return SieveOfEratosthenes(until)
}

// SieveOfEratosthenes is a simple, ancient algorithm for finding all prime numbers up to any given limit.
// https://en.wikipedia.org/wiki/Sieve_of_Eratosthenes#Pseudocode
func SieveOfEratosthenes(until int) []int {

	// 1. Create a list of consecutive integers from 2 through n: (2, 3, 4, ..., n).
	// ie. [number] = isPrime
	// eg. [13] = true
	// note: golang defaults bool to false, so we'll actually do the inverse (eg. [13] = false)
	numbers := make([]bool, until+1)

	// estimate how high we need to go to get a prime close to the until value
	// √n rounded up
	end := int(math.Sqrt(float64(until)) + 0.5)

	// iterate over all numbers
	// 2. Initially, let p equal 2, the smallest prime number.
	for p := 2; p <= end; p++ {

		// 4. Find the first number greater than p in the list that is not marked.
		//    If there was no such number, stop.
		//    Otherwise, let p now equal this new number (which is the next prime), and repeat from step 3.
		// ie. if the number is already "crossed off", it's divisable by a smaller number,
		// so there's no point in even using it here
		// eg. don't check where p=4, because 4 is divisable by 2, so we've already checked off 2,4,6,8,10,12,etc
		if numbers[p] == false {

			// 3. Enumerate the multiples of p by counting in increments of p from 2p to n, and mark them in the list
			//    (these will be 2p, 3p, 4p, ...; the p itself should not be marked).
			j := 0
			for h := 0; ; h++ {
				j = p*p + h*p
				if j > until {
					break
				}
				numbers[j] = true
			}
		}
	}
	// 5. When the algorithm terminates,
	//    the numbers remaining not marked in the list are all the primes below n.

	primes = []int{}
	for i := 2; i < len(numbers); i++ {
		if numbers[i] == false {
			primes = append(primes, i)
		}
	}

	return primes
}

// Generator returns a generator which returns the next prime in sequence
// (brute force method)
func Generator() func() int {

	p := 1

	return func() int {

		for {

			p++

			// test all integers below p
			end := int(math.Sqrt(float64(p)) + 0.5)
			for j := end; j > 0; j-- {
				if p%j == 0 {

					if j == 1 {
						// evenly divisable by 1, prime!
						return p
					}

					// evenly divisable by not 1, not prime!
					break
				}
			}
		}
	}
}

func IsPrime(n int) bool {

	// check cache
	if len(primes) > 0 && primes[len(primes)-1] > n {
		for _, v := range primes {
			if v == n {
				return true
			}
		}
		return false
	}

	// test all integers below n
	end := int(math.Sqrt(float64(n)) + 0.5)
	for j := end; j > 0; j-- {
		if n%j == 0 {

			if j == 1 {
				// evenly divisable by 1, prime!
				return true
			}

			// evenly divisable by not 1, not prime!
			break
		}
	}

	return false

}
