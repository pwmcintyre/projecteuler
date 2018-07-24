package numbers

import (
	"reflect"
	"testing"
)

func TestAbundantNumbers(t *testing.T) {

	var examples = []struct {
		input  int
		answer []int
	}{
		{10, []int{}},
		{20, []int{12, 18}},
		{120, []int{12, 18, 20, 24, 30, 36, 40, 42, 48, 54, 56, 60, 66, 70, 72, 78, 80, 84, 88, 90, 96, 100, 102, 104, 108, 112, 114}},
	}

	for _, ex := range examples {
		result := AbundantNumbers(ex.input)
		if !reflect.DeepEqual(ex.answer, result) {
			t.Errorf("\nresult: %+v\nexpected: %d\ninput: %+v", result, ex.answer, ex.input)
		}
	}
}

func TestIsAbundantNumber(t *testing.T) {

	var examples = []struct {
		input  int
		answer bool
	}{
		{5, false},
		{12, true},
		{120, true},
	}

	for _, ex := range examples {
		result := IsAbundantNumber(ex.input)
		if !reflect.DeepEqual(ex.answer, result) {
			t.Errorf("\nresult: %+v\nexpected: %v\ninput: %+v", result, ex.answer, ex.input)
		}
	}
}

func TestNotSumOfAbundantNumbers(t *testing.T) {

	var examples = []struct {
		input  int
		answer []int
	}{
		{26, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 25}},
	}

	for _, ex := range examples {
		result := NotSumOfAbundantNumbers(ex.input)
		if !reflect.DeepEqual(ex.answer, result) {
			t.Errorf("\nresult: %+v\nexpected: %d\ninput: %+v", result, ex.answer, ex.input)
		}
	}
}
