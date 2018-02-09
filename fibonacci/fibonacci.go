package fibonacci

import "math/big"

// Fibonacci Length returns a number a digits in the Fibonacci sequence
func Length(length int) []int {

	sequence := []int{1, 2}

	for len(sequence) < length {
		i := len(sequence)
		next := sequence[i-1] + sequence[i-2]
		sequence = append(sequence, next)
	}

	return sequence
}

// Fibonacci Until return numbers of the Fibonacci sequence less than or equal to a given maximum
func Until(max int) []int {

	sequence := []int{1, 2}

	for {
		i := len(sequence)
		next := sequence[i-1] + sequence[i-2]
		if next > max {
			break
		}
		sequence = append(sequence, next)
	}

	return sequence
}

// Generator returns the next permutation of the fibonacci sequence
func Generator() func() int {

	n1 := int(0)
	n2 := int(1)

	return func() (n int) {
		n = n1 + n2
		n1 = n2
		n2 = n
		return
	}
}

// Generator returns the next permutation of the fibonacci sequence
func BigGenerator() func() *big.Int {

	n1 := big.NewInt(0)
	n2 := big.NewInt(1)

	return func() (n *big.Int) {
		n = n1.Add(n1, n2)
		n1 = n2
		n2 = n
		return
	}
}
