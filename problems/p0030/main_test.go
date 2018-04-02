package main

import (
	"reflect"
	"testing"
)

func TestRun(t *testing.T) {

	var examples = []struct {
		below  int
		power  int
		answer int
	}{
		{10000, 4, 19316},
		{100000000, 5, 443839},
	}

	for _, ex := range examples {
		result := Run(ex.below, ex.power)
		if !reflect.DeepEqual(ex.answer, result) {
			t.Errorf("\nresult: %+v\nexpected: %d\nbelow: %+v", result, ex.answer, ex.below)
		}
	}
}

func TestSumOfPower(t *testing.T) {

	var examples = []struct {
		below  int
		power  int
		answer []int
	}{
		{10000, 4, []int{1634, 8208, 9474}},
	}

	for _, ex := range examples {
		result := SumOfPower(ex.below, ex.power)
		if !reflect.DeepEqual(ex.answer, result) {
			t.Errorf("\nresult: %+v\nexpected: %d\nbelow: %+v", result, ex.answer, ex.below)
		}
	}
}

func TestIsSumOfPower(t *testing.T) {

	var examples = []struct {
		input  int
		power  int
		answer bool
	}{
		{1, 1, true},
		{2, 2, false},
		{1634, 4, true},
	}

	for _, ex := range examples {
		result := IsSumOfPower(ex.input, ex.power)
		if !reflect.DeepEqual(ex.answer, result) {
			t.Errorf("\nresult: %+v\nexpected: %+v\nbelow: %+v", result, ex.answer, ex.input)
		}
	}
}
