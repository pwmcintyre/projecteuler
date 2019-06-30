package prime

import (
	prime "github.com/pwmcintyre/projecteuler/prime/IncrementalSieveOfEratosthenes"
)

type Prime struct {
	cache []int
	next  <-chan int
}

func New() *Prime {
	return &Prime{
		cache: []int{1},
		next:  prime.Generator(),
	}
}

func (p *Prime) increment() {

	p.cache = append(p.cache, <-p.next)

}

func (p *Prime) Is(n int) bool {

	// ensure we have enough known primes
	for n > p.cache[len(p.cache)-1] {

		// generate next prime
		p.increment()

	}

	// check n in list of primes
	for _, prime := range p.cache {
		if prime == n {
			return true
		}
	}

	// not prime
	return false

}
