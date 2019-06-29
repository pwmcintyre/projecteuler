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
		{1, 23},      // 23
		{2, 60},      // 23 + 37
		{11, 748317}, //  23, 23, 37, 23, 37, 53, 73, 313, 317, 373, 797, 3137, 3797, 739397
	}

	for _, ex := range examples {
		result := Run(ex.below)
		if !reflect.DeepEqual(ex.answer, result) {
			t.Errorf("\nresult: %+v\nexpected: %d\nbelow: %+v", result, ex.answer, ex.below)
		}
	}
}

func TestIsTruncatablePrime(t *testing.T) {

	var examples = []struct {
		n      int
		answer bool
	}{
		{3797, true},
		{23, true},
		{53, true},
		{3, false},  // less than 10 is not considered truncatable
		{11, false}, // 1 is not considered prime
		{20, false}, // 20 not prime
		{83, false}, // 8 not prime
	}

	for _, ex := range examples {
		result := IsTruncatablePrime(ex.n)
		if !reflect.DeepEqual(ex.answer, result) {
			t.Errorf("\nresult: %+v\nexpected: %v\nn: %+v", result, ex.answer, ex.n)
		}
	}
}
