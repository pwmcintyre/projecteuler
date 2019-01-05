package main

import (
	"reflect"
	"testing"
)

func TestConsecutivePrimes(t *testing.T) {

	var examples = []struct {
		input  input
		answer answer
	}{
		{input{1, 41}, 40},
	}

	for _, ex := range examples {
		result := ConsecutivePrimes(ex.input)
		if !reflect.DeepEqual(ex.answer, result) {
			t.Errorf("\nresult: %+v\nexpected: %d\ninput: %+v", result, ex.answer, ex.input)
		}
	}
}
