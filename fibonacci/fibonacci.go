package fibonacci

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
