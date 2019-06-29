package functions

import (
	"strconv"

	"github.com/pwmcintyre/projecteuler/arrays"
)

func TruncationsOf(n int) []int {

	// add original number
	truncations := []int{n}

	// convert to string for truncation
	nstr := strconv.Itoa(n)

	// left to right
	for i := 1; i < len(nstr); i++ {
		n, _ := strconv.Atoi(nstr[i:])
		truncations = append(truncations, n)
	}

	// right to left
	for i := 1; i < len(nstr); i++ {
		n, _ := strconv.Atoi(nstr[:i])
		truncations = append(truncations, n)
	}

	return arrays.UniqueInt(truncations)

}
