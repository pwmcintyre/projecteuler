package main

import (
	"reflect"
	"testing"
)

func TestRun(t *testing.T) {

	var examples = []struct {
		below  int
		prec   int
		answer int
	}{
		{2, 2, 0},
		{10, 20, 7},
		{25, 100, 23},
	}

	for _, ex := range examples {
		result := Run(ex.below)
		if !reflect.DeepEqual(ex.answer, result) {
			t.Errorf("\nresult: %+v\nexpected: %d\nbelow: %+v", result, ex.answer, ex.below)
		}
	}
}

func TestGetCycle(t *testing.T) {

	var examples = []struct {
		numerator   int
		denominator int
		start       int
		str         string
		cycle       string
	}{
		{1, 1, 0, "1.0", ""},
		{5, 3, 0, "1.(6)", "6"},
		{2, 1, 0, "2.0", ""},
		{2, 3, 0, "0.(6)", "6"},
		{1, 2, 0, "0.5", ""},
		{1, 3, 0, "0.(3)", "3"},
		{1, 4, 0, "0.25", ""},
		{1, 5, 0, "0.2", ""},
		{1, 6, 1, "0.1(6)", "6"},
		{1, 7, 0, "0.(142857)", "142857"},
		{1, 8, 0, "0.125", ""},
		{1, 9, 0, "0.(1)", "1"},
		{1, 10, 0, "0.1", ""},
		{1, 23, 0, "0.(0434782608695652173913)", "0434782608695652173913"},
	}

	for _, ex := range examples {
		_, _, cycle := GetCycle(ex.numerator, ex.denominator, 50)
		// if !reflect.DeepEqual(ex.start, start) {
		// 	t.Errorf("\nstart: %+v\nexpected: %d\ninput: %+v / %+v", start, ex.start, ex.numerator, ex.denominator)
		// }
		if !reflect.DeepEqual(ex.cycle, cycle) {
			t.Errorf("\ncycle: %s\nexpected: %s\ninput: %+v / %+v", cycle, ex.cycle, ex.numerator, ex.denominator)
		}
		// if !reflect.DeepEqual(ex.str, str) {
		// 	t.Errorf("\nstr: %s\nexpected: %s\ninput: %+v / %+v", str, ex.str, ex.numerator, ex.denominator)
		// }
	}
}
