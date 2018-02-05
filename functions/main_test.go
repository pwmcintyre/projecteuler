package functions

import (
	"reflect"
	"testing"
)

func TestGCD(t *testing.T) {

	// 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...
	examples := []struct {
		input  []int
		answer int
	}{
		{[]int{48, 18}, 6},
		{[]int{10, 20}, 10},
		{[]int{10, 21}, 1},
		{[]int{123, 321}, 3},
	}

	for _, ex := range examples {
		result := GCD(ex.input[0], ex.input[1])
		if !reflect.DeepEqual(ex.answer, result) {
			t.Errorf("\nresult: %+v\nexpected: %d\ninput: %+v", result, ex.answer, ex.input)
		}
	}
}

func TestIsPalindromic(t *testing.T) {

	examples := []struct {
		input  int
		answer bool
	}{
		{123, false},
		{121, true},
		{10000000000000001, true},
	}

	for _, ex := range examples {
		result := IsPalindromic(ex.input)
		if !reflect.DeepEqual(ex.answer, result) {
			t.Errorf("\nresult: %+v\nexpected: %d\ninput: %+v", result, ex.answer, ex.input)
		}
	}
}

func TestMaxPositiveInt(t *testing.T) {

	examples := []struct {
		input  []int
		answer int
	}{
		{[]int{1, 2, 3, 4}, 4},
		{[]int{4, 1, 2}, 4},
		{[]int{1}, 1},
	}

	for _, ex := range examples {
		result := MaxPositiveInt(ex.input)
		if !reflect.DeepEqual(ex.answer, result) {
			t.Errorf("\nresult: %+v\nexpected: %d\ninput: %+v", result, ex.answer, ex.input)
		}
	}
}

func TestDivisors(t *testing.T) {

	examples := []struct {
		input  int
		answer []int
	}{
		{1, []int{1}},
		{3, []int{1, 3}},
		{6, []int{1, 2, 3, 6}},
		{10, []int{1, 2, 5, 10}},
		{15, []int{1, 3, 5, 15}},
		{21, []int{1, 3, 7, 21}},
		{28, []int{1, 2, 4, 7, 14, 28}},
		{1275, []int{1, 3, 5, 15, 17, 25, 51, 75, 85, 255, 425, 1275}},
	}

	for _, ex := range examples {
		result := Divisors(ex.input)
		if !reflect.DeepEqual(ex.answer, result) {
			t.Errorf("\nresult: %+v\nexpected: %d\ninput: %+v", result, ex.answer, ex.input)
		}
	}
}

func TestDivisorCount(t *testing.T) {

	examples := []struct {
		input  int
		answer int
	}{
		{1, 1},
		{3, 2},
		{6, 4},
		{10, 4},
		{28, 6},
		{1275, 12},
	}

	for _, ex := range examples {
		result := DivisorCount(ex.input)
		if !reflect.DeepEqual(ex.answer, result) {
			t.Errorf("\nresult: %+v\nexpected: %d\ninput: %+v", result, ex.answer, ex.input)
		}
	}
}

func TestInArray(t *testing.T) {

	examples := []struct {
		array  []string
		find   string
		answer int
	}{
		{[]string{"Mary", "Anna", "Beth", "Johnny", "Beth"}, "Anna", 1},
	}

	for _, ex := range examples {
		_, result := InArray(ex.find, ex.array)
		if !reflect.DeepEqual(ex.answer, result) {
			t.Errorf("\nresult: %+v\nexpected: %d\narray: %+v", result, ex.answer, ex.array)
		}
	}
}
