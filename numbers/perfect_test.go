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
		{300, []int{220, 284}},
		{2000, []int{220, 284, 1184, 1210}},
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
		{200, false},
		{220, true},
		{1210, true},
	}

	for _, ex := range examples {
		result := IsPerfectNumber(ex.input)
		if !reflect.DeepEqual(ex.answer, result) {
			t.Errorf("\nresult: %+v\nexpected: %d\ninput: %+v", result, ex.answer, ex.input)
		}
	}
}
