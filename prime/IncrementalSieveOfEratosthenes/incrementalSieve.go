package prime

type IPrimeInt interface {
	Is(int) bool
}

type PrimeInt struct {
	composites map[int][]int
	primes     []int
	out        chan int
	last       int
}

func New() IPrimeInt {
	p := PrimeInt{}
	p.primes = make([]int, 100)
	return &p
}

// generates another prime
func (p PrimeInt) gen() int {

	p.last = <-p.out
	return p.last

}

func (p PrimeInt) Is(n int) bool {

	// ensure we have enough known primes
	for n > p.last {

		// generate next prime
		p.gen()

		// save it
		p.primes = append(p.primes, p.last)

	}

	// check n in list of primes
	for _, prime := range p.primes {
		if prime == n {
			return true
		}
	}

	// not prime
	return false

}
