package divisors

import (
	"reflect"
	"testing"
)

func TestDivisors(t *testing.T) {

	examples := []struct {
		input  int
		answer []int
	}{
		{1, []int{1}},
		{3, []int{1, 3}},
		{6, []int{1, 2, 3, 6}},
		{10, []int{1, 2, 5, 10}},
		{15, []int{1, 3, 5, 15}},
		{21, []int{1, 3, 7, 21}},
		{28, []int{1, 2, 4, 7, 14, 28}},
		{1275, []int{1, 3, 5, 15, 17, 25, 51, 75, 85, 255, 425, 1275}},
	}

	for _, ex := range examples {
		result := Divisors(ex.input)
		if !reflect.DeepEqual(ex.answer, result) {
			t.Errorf("\nresult: %+v\nexpected: %d\ninput: %+v", result, ex.answer, ex.input)
		}
	}
}

func TestProperDivisors(t *testing.T) {

	examples := []struct {
		input  int
		answer []int
	}{
		{1, []int{}},
		{3, []int{1}},
		{6, []int{1, 2, 3}},
		{10, []int{1, 2, 5}},
		{15, []int{1, 3, 5}},
		{21, []int{1, 3, 7}},
		{28, []int{1, 2, 4, 7, 14}},
		{1275, []int{1, 3, 5, 15, 17, 25, 51, 75, 85, 255, 425}},
	}

	for _, ex := range examples {
		result := ProperDivisors(ex.input)
		if !reflect.DeepEqual(ex.answer, result) {
			t.Errorf("\nresult: %+v\nexpected: %d\ninput: %+v", result, ex.answer, ex.input)
		}
	}
}

func TestDivisorCount(t *testing.T) {

	examples := []struct {
		input  int
		answer int
	}{
		{1, 1},
		{3, 2},
		{6, 4},
		{10, 4},
		{28, 6},
		{1275, 12},
	}

	for _, ex := range examples {
		result := DivisorCount(ex.input)
		if !reflect.DeepEqual(ex.answer, result) {
			t.Errorf("\nresult: %+v\nexpected: %d\ninput: %+v", result, ex.answer, ex.input)
		}
	}
}
