// package main
/*

Euler discovered the remarkable quadratic formula:

n2+n+41
It turns out that the formula will produce 40 primes for the consecutive integer values 0≤n≤39. However, when n=40,402+40+41=40(40+1)+41 is divisible by 41, and certainly when n=41,412+41+41 is clearly divisible by 41.

The incredible formula n2−79n+1601 was discovered, which produces 80 primes for the consecutive values 0≤n≤79. The product of the coefficients, −79 and 1601, is −126479.

Considering quadratics of the form:

n2+an+b, where |a|<1000 and |b|≤1000

where |n| is the modulus/absolute value of n
e.g. |11|=11 and |−4|=4
Find the product of the coefficients, a and b, for the quadratic expression that produces the maximum number of primes for consecutive values of n, starting with n=0.

*/

package main

import (
	"fmt"
	"math"

	"github.com/pwmcintyre/projecteuler/prime"

	"github.com/sirupsen/logrus"
)

type input = struct{ a, b int }
type answer = int

// Run run
func Run(in input) answer {

	// logger := logrus.StandardLogger()

	// pre-compute primes cache for performance
	prime.SieveOfEratosthenes(quadratic(in.a, in.b, 10000))

	//
	// var bestInput input
	var bestResult int

	maxa := int(math.Abs(float64(in.a)))
	maxb := int(math.Abs(float64(in.b)))

	for a := -maxa; a <= maxa; a++ {
		for b := -maxb; b <= maxb; b++ {

			v := ConsecutivePrimes(a, b)

			// logger.WithField("input", input{a, b}).WithField("result", v).Debug("result")

			if v > bestResult {
				bestResult = v
				// bestInput = input{a, b}
			}
		}
	}

	// logger.WithField("input", bestInput).WithField("result", bestResult).Info("best")

	return bestResult

}

// n^2+an+b
func ConsecutivePrimes(a, b int) int {

	consecutivePrimes := 0
	for n := 0; ; n++ {

		v := quadratic(a, b, n)

		// check if v is prime
		isPrime := prime.IsPrime(v)
		if isPrime {
			consecutivePrimes++
			continue
		}

		break

	}

	return consecutivePrimes

}

// n^2+an+b
func quadratic(a, b, n int) int {
	return int(math.Pow(float64(n), 2)) + a*n + b
}

func main() {
	logrus.StandardLogger().SetLevel(logrus.DebugLevel)
	// fmt.Println(Run(input{1, 41}))
	// fmt.Println(Run(input{-79, 1601}))
	// fmt.Println(Run(input{1000, 1000}))
	fmt.Println(GoRun(input{1000, 1000}, 4))
}
