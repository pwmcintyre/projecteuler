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
		{10, 4},
		{100, 13},
		{200, 17},
		{1000000, 55},
	}

	for _, ex := range examples {
		result := Run(ex.below)
		if !reflect.DeepEqual(ex.answer, result) {
			t.Errorf("\nresult: %+v\nexpected: %d\nbelow: %+v", result, ex.answer, ex.below)
		}
	}
}

func TestIsCircularPrime(t *testing.T) {

	var examples = []struct {
		n               int
		isCircularPrime bool
	}{
		{2, true},
		{3, true},
		{5, true},
		{7, true},
		{11, true},
		{13, true},
		{17, true},
		{31, true},
		{37, true},
		{71, true},
		{73, true},
		{79, true},
		{97, true},
		{197, true},
		{4, false},
		{10, false},
		{98, false},
		{99, false},
		{100, false},
	}

	for _, ex := range examples {
		result := IsCircularPrime(ex.n)
		if !reflect.DeepEqual(ex.isCircularPrime, result) {
			t.Errorf("\nresult: %+v\nexpected: %v\nn: %+v", result, ex.isCircularPrime, ex.n)
		}
	}
}

// BenchmarkIsCircularPrime-4   	 1000000	      5080 ns/op	     147 B/op	       7 allocs/op
func BenchmarkIsCircularPrime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsCircularPrime(i)
	}
}
