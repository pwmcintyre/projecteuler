package main

import (
	"reflect"
	"testing"
)

func TestRun(t *testing.T) {

	var examples = []struct {
		input  int
		index  int
		answer string
	}{
		{2, 0, "01"},
		{3, 0, "012"},
		{3, 5, "210"},
		{10, 0, "0123456789"},
		{10, 1, "0123456798"},
		{10, 1000000 - 1, "2783915460"},
	}

	for _, ex := range examples {
		result := Run(ex.input, ex.index)
		if !reflect.DeepEqual(ex.answer, result) {
			t.Errorf("\nresult: %+v\nexpected: %d\ninput: %+v", result, ex.answer, ex.input)
		}
	}
}

func TestNthPermutationOfOrderedArrangement(t *testing.T) {

	var examples = []struct {
		input  []int
		index  int
		answer string
	}{
		// 01 10
		{[]int{0, 1}, 0, "01"},
		{[]int{0, 1}, 1, "10"},
		// 012   021   102   120   201   210
		{[]int{0, 1, 2}, 0, "012"},
		{[]int{0, 1, 2}, 1, "021"},
		{[]int{0, 1, 2}, 2, "102"},
		{[]int{0, 1, 2}, 3, "120"},
		{[]int{0, 1, 2}, 4, "201"},
		{[]int{0, 1, 2}, 5, "210"},
	}

	for _, ex := range examples {
		result := NthPermutationOfOrderedArrangement(ex.input, ex.index)
		if !reflect.DeepEqual(ex.answer, result) {
			t.Errorf("\nresult: %+v\nexpected: %d\ninput: %+v", result, ex.answer, ex.input)
		}
	}
}

func TestGenDigits(t *testing.T) {

	var examples = []struct {
		input  int
		answer []int
	}{
		{3, []int{0, 1, 2}},
	}

	for _, ex := range examples {
		result := GenDigits(ex.input)
		if !reflect.DeepEqual(ex.answer, result) {
			t.Errorf("\nresult: %+v\nexpected: %d\ninput: %+v", result, ex.answer, ex.input)
		}
	}
}
