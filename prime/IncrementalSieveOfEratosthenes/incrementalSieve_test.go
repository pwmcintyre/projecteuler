package prime

import (
	"reflect"
	"testing"
)

// first 46 primes
var testPrimes = []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103, 107, 109, 113, 127, 131, 137, 139, 149, 151, 157, 163, 167, 173, 179, 181, 191, 193, 197, 199}

func TestNext(t *testing.T) {

	primes := next()

	for i, ex := range testPrimes {
		result := <-primes
		if !reflect.DeepEqual(ex, result) {
			t.Errorf("\nresult: %+v\nexpected: %d\nindex: %+v", result, ex, i)
		}
		t.Log(result)
	}
}

func TestIsPrime(t *testing.T) {

	examples := []struct {
		input  int
		answer bool
	}{
		{3, true},
		{4, false},
		{5, true},
		{6, false},
		{199, true},
		{1000, false},
	}

	for _, ex := range examples {

		prime := New()

		result := prime.Is(ex.input)
		if !reflect.DeepEqual(ex.answer, result) {
			t.Errorf("\nresult: %+v\nexpected: %v\ninput: %+v", result, ex.answer, ex.input)
		}
		t.Log(result)
	}
}
