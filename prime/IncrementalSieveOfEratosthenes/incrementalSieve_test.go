package prime

import (
	"reflect"
	"testing"
)

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
	}
}
