package main

import (
	"reflect"
	"testing"
)

var examples = []struct {
	input   int
	factors []int
	answer  int
}{
	{100, []int{2, 2, 5, 5}, 5},
	{144, []int{2, 2, 2, 2, 3, 3}, 3},
	{315, []int{3, 3, 5, 7}, 7},
	{360, []int{2, 2, 2, 3, 3, 5}, 5},
	{8051, []int{83, 97}, 97},
	{13195, []int{5, 7, 13, 29}, 29},
	{600851475143, []int{71, 839, 1471, 6857}, 6857},
}

func TestPrimeFactors(t *testing.T) {

	for _, ex := range examples {
		result := PrimeFactors(ex.input)
		if !reflect.DeepEqual(ex.factors, result) {
			t.Errorf("\nresult: %+v\nexpected: %d\ninput: %+v", result, ex.answer, ex.input)
		}
	}
}

func TestRun(t *testing.T) {

	for _, ex := range examples {
		result := Run(ex.input)
		if !reflect.DeepEqual(ex.answer, result) {
			t.Errorf("\nresult: %+v\nexpected: %d\ninput: %+v", result, ex.answer, ex.input)
		}
	}
}
