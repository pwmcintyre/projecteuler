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
	"math"

	"github.com/pwmcintyre/projecteuler/prime"
	"github.com/sirupsen/logrus"
)

type batch = struct{ mina, maxa, minb, maxb int }

func worker(id int, jobs <-chan batch, results chan<- answer) {

	logger := logrus.StandardLogger().WithField("id", id)
	logger.Info("worker started")

	for j := range jobs {

		logger = logger.WithField("job", j)
		logger.Info("job started")

		result := RunBatch(j)
		results <- result

		logger.WithField("result", result).Info("job ended")

	}
}

func GoRun(in input, workerCount int) answer {

	// pre-compute primes cache for performance
	prime.SieveOfEratosthenes(quadratic(in.a, in.b, 10000))

	jobs := make(chan batch, 100)
	results := make(chan answer, 100)

	// make some workers
	for w := 0; w < workerCount; w++ {
		go worker(w, jobs, results)
	}

	// make some jobs
	// we'll make 1 batch = a since A to iterate over all possible B's
	maxa := int(math.Abs(float64(in.a)))
	maxb := int(math.Abs(float64(in.b)))
	jobCount := 0
	for a := -maxa; a <= maxa; a++ {
		batch := batch{a, a, -maxb, maxb}
		jobs <- batch
		jobCount++
	}
	close(jobs)

	// collect results
	var bestResult int
	for j := 0; j < jobCount; j++ {
		r := <-results
		if r > bestResult {
			bestResult = r
		}
	}

	return bestResult

}

// Run run
func RunBatch(in batch) answer {

	logger := logrus.StandardLogger()

	//
	var bestInput input
	var bestResult int

	for a := in.mina; a <= in.maxa; a++ {
		for b := in.minb; b <= in.maxb; b++ {

			v := ConsecutivePrimes(a, b)

			if v > bestResult {
				bestResult = v
				bestInput = input{a, b}
			}
		}
	}

	logger.WithField("input", bestInput).WithField("result", bestResult).Info("batch best")

	return bestResult

}
