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

// NextPrime is an excellent implementation using the Incremental Sieve of Eratosthenes method
// http://lukelafountaine.com/programming/math/2016/12/10/Incremental-Sieve.html
func next() <-chan int {

	out := make(chan int, 1)

	go func() {

		out <- 2
		num := 3
		composites := make(map[int][]int)

		for {
			if _, ok := composites[num]; !ok {
				out <- num
				composites[num*num] = []int{num}
			} else {
				for _, prime := range composites[num] {
					next := num + prime
					for next%2 == 0 {
						next += prime
					}
					if _, ok := composites[next]; ok {
						composites[next] = append(composites[next], prime)
					} else {
						composites[next] = []int{prime}
					}
				}
				delete(composites, num)
			}
			num += 2
		}
	}()

	return out

}
