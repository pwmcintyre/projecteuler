package main

import (
	"reflect"
	"testing"
)

func TestGoRun(t *testing.T) {

	var examples = []struct {
		input  input
		answer answer
		count  int
	}{
		// {input{1, 10}, 5, 2},
		{input{-79, 1601}, 80, 1},
		// {input{-79, 1601}, 80, 2},
		// {input{-79, 1601}, 80, 10},
	}

	for _, ex := range examples {
		result := GoRun(ex.input, ex.count)
		if !reflect.DeepEqual(ex.answer, result) {
			t.Errorf("\nresult: %+v\nexpected: %d\ninput: %+v", result, ex.answer, ex.input)
		}
	}
}

func TestRun(t *testing.T) {

	var examples = []struct {
		input  input
		answer answer
	}{
		{input{-79, 1601}, 80},
	}

	for _, ex := range examples {
		result := Run(ex.input)
		if !reflect.DeepEqual(ex.answer, result) {
			t.Errorf("\nresult: %+v\nexpected: %d\ninput: %+v", result, ex.answer, ex.input)
		}
	}
}

func TestConsecutivePrimes(t *testing.T) {

	var examples = []struct {
		a, b   int
		answer int
	}{
		{1, 41, 40},
		{-79, 1601, 80},
	}

	for _, ex := range examples {
		result := ConsecutivePrimes(ex.a, ex.b)
		if !reflect.DeepEqual(ex.answer, result) {
			t.Errorf("\nresult: %+v\nexpected: %d\ninput: %+v", result, ex.answer, struct{ a, b int }{ex.a, ex.b})
		}
	}
}
