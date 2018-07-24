package numbers

import (
	"reflect"
	"testing"
)

func TestIsPalindromic(t *testing.T) {

	examples := []struct {
		input  int
		answer bool
	}{
		{123, false},
		{121, true},
		{10000000000000001, true},
	}

	for _, ex := range examples {
		result := IsPalindromic(ex.input)
		if !reflect.DeepEqual(ex.answer, result) {
			t.Errorf("\nresult: %+v\nexpected: %v\ninput: %+v", result, ex.answer, ex.input)
		}
	}
}
