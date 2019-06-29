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
		{2200, 2143},
		{3000, 2341},
		{987654321, 2341},
	}

	for _, ex := range examples {
		result := Run(ex.below)
		if !reflect.DeepEqual(ex.answer, result) {
			t.Errorf("\nresult: %+v\nexpected: %d\nbelow: %+v", result, ex.answer, ex.below)
		}
	}
}

func TestIsPandigital(t *testing.T) {

	var examples = []struct {
		n      int
		answer bool
	}{
		{2143, true},
		{123456789, true},
		{1, true},
		{12, true},
		{13, false},
		{34, false},
	}

	for _, ex := range examples {
		result := IsPandigital(ex.n)
		if !reflect.DeepEqual(ex.answer, result) {
			t.Errorf("\nresult: %+v\nexpected: %v\nn: %+v", result, ex.answer, ex.n)
		}
	}
}
