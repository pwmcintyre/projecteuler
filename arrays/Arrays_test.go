package arrays

import (
	"reflect"
	"testing"
)

func TestMaxPositiveInt(t *testing.T) {

	examples := []struct {
		input  []int
		answer int
	}{
		{[]int{1, 2, 3, 4}, 4},
		{[]int{4, 1, 2}, 4},
		{[]int{1}, 1},
	}

	for _, ex := range examples {
		result := MaxPositiveInt(ex.input)
		if !reflect.DeepEqual(ex.answer, result) {
			t.Errorf("\nresult: %+v\nexpected: %d\ninput: %+v", result, ex.answer, ex.input)
		}
	}
}

func TestInArray(t *testing.T) {

	examples := []struct {
		array  []string
		find   string
		answer int
	}{
		{[]string{"Mary", "Anna", "Beth", "Johnny", "Beth"}, "Anna", 1},
	}

	for _, ex := range examples {
		_, result := InArray(ex.find, ex.array)
		if !reflect.DeepEqual(ex.answer, result) {
			t.Errorf("\nresult: %+v\nexpected: %d\narray: %+v", result, ex.answer, ex.array)
		}
	}
}

func TestUniqueInt(t *testing.T) {

	examples := []struct {
		input  []int
		output []int
	}{
		{[]int{1, 2, 3, 4}, []int{1, 2, 3, 4}},
		{[]int{1, 1, 1, 4}, []int{1, 4}},
	}

	for _, ex := range examples {
		result := UniqueInt(ex.input)
		if !reflect.DeepEqual(ex.output, result) {
			t.Errorf("\nresult: %+v\nexpected: %d\ninput: %+v", result, ex.output, ex.input)
		}
	}
}
