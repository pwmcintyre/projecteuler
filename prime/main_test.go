package prime

import (
	"reflect"
	"testing"
)

// first 46 primes
var primes = []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103, 107, 109, 113, 127, 131, 137, 139, 149, 151, 157, 163, 167, 173, 179, 181, 191, 193, 197, 199}

func TestUntil(t *testing.T) {

	examples := []struct {
		input  int
		answer []int
	}{
		{1, primes[0:0]},
		{2, primes[0:1]},
		{10, primes[0:4]},
		{20, primes[0:8]},
	}

	for _, ex := range examples {
		result := Until(ex.input)
		if !reflect.DeepEqual(ex.answer, result) {
			t.Errorf("\nresult: %+v\nexpected: %d\ninput: %+v", result, ex.answer, ex.input)
		}
	}
}
