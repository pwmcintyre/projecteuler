package main

import (
	"reflect"
	"testing"
)

var examples = []struct {
	input  int
	answer int
}{
	{13, 9},
	{100, 97},
	{1000000, 837799},
}

func TestRun(t *testing.T) {
	for _, ex := range examples {
		result := Run(ex.input)
		if !reflect.DeepEqual(ex.answer, result) {
			t.Errorf("\nresult: %+v\nexpected: %d\ninput: %+v", result, ex.answer, ex.input)
		}
	}
}

func TestWalk(t *testing.T) {
	var examples = []struct {
		input  int
		answer int
	}{
		{1, 0},
		{2, 1},
		{13, 9},
		// {100, 1},
		// {1000000, 1},
	}
	for _, ex := range examples {
		result := Walk(ex.input)
		if !reflect.DeepEqual(ex.answer, result) {
			t.Errorf("\nresult: %+v\nexpected: %d\ninput: %+v", result, ex.answer, ex.input)
		}
	}
}
