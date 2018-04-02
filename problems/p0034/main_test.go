package main

import (
	"reflect"
	"testing"
)

func TestRun(t *testing.T) {

	var examples = []struct {
		below  int
		answer int
	}{
		{150, 145},
		{100000, 40730},
		{100000000, 40730},
	}

	for _, ex := range examples {
		result := Run(ex.below)
		if !reflect.DeepEqual(ex.answer, result) {
			t.Errorf("\nresult: %+v\nexpected: %d\nbelow: %+v", result, ex.answer, ex.below)
		}
	}
}

func TestSumOfFactorials(t *testing.T) {

	var examples = []struct {
		below  int
		answer []int
	}{
		{150, []int{145}},
		{100000, []int{145, 40585}},
	}

	for _, ex := range examples {
		result := SumOfFactorials(ex.below)
		if !reflect.DeepEqual(ex.answer, result) {
			t.Errorf("\nresult: %+v\nexpected: %d\nbelow: %+v", result, ex.answer, ex.below)
		}
	}
}

func TestIsSumOfFactorials(t *testing.T) {

	var examples = []struct {
		input  int
		answer bool
	}{
		{1, true},
		{2, true},
		{3, false},
		{145, true},
		{146, false},
	}

	for _, ex := range examples {
		result := IsSumOfFactorials(ex.input)
		if !reflect.DeepEqual(ex.answer, result) {
			t.Errorf("\nresult: %+v\nexpected: %+v\nbelow: %+v", result, ex.answer, ex.input)
		}
	}
}
