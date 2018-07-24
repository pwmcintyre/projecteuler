package numbers

import (
	"reflect"
	"testing"
)

func TestAmicableNumbers(t *testing.T) {

	var examples = []struct {
		input  int
		answer []int
	}{
		{300, []int{220, 284}},
		{2000, []int{220, 284, 1184, 1210}},
	}

	for _, ex := range examples {
		result := AmicableNumbers(ex.input)
		if !reflect.DeepEqual(ex.answer, result) {
			t.Errorf("\nresult: %+v\nexpected: %d\ninput: %+v", result, ex.answer, ex.input)
		}
	}
}

func TestIsAmicableNumber(t *testing.T) {

	var examples = []struct {
		input  int
		answer bool
	}{
		{200, false},
		{220, true},
		{1210, true},
	}

	for _, ex := range examples {
		result := IsAmicableNumber(ex.input)
		if !reflect.DeepEqual(ex.answer, result) {
			t.Errorf("\nresult: %+v\nexpected: %v\ninput: %+v", result, ex.answer, ex.input)
		}
	}
}
