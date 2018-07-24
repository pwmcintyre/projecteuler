package numbers

import (
	"reflect"
	"testing"
)

func TestPerfectNumbers(t *testing.T) {

	var examples = []struct {
		input  int
		answer []int
	}{
		{10, []int{6}},
		{30, []int{6, 28}},
	}

	for _, ex := range examples {
		result := PerfectNumbers(ex.input)
		if !reflect.DeepEqual(ex.answer, result) {
			t.Errorf("\nresult: %+v\nexpected: %d\ninput: %+v", result, ex.answer, ex.input)
		}
	}
}

func TestIsPerfectNumber(t *testing.T) {

	var examples = []struct {
		input  int
		answer bool
	}{
		{5, false},
		{6, true},
		{28, true},
	}

	for _, ex := range examples {
		result := IsPerfectNumber(ex.input)
		if !reflect.DeepEqual(ex.answer, result) {
			t.Errorf("\nresult: %+v\nexpected: %v\ninput: %+v", result, ex.answer, ex.input)
		}
	}
}
