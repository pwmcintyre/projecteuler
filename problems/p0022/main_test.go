package main

import (
	"io/ioutil"
	"reflect"
	"testing"
)

func TestRun(t *testing.T) {

	var examples = []struct {
		input  string
		answer int
	}{
		{`"ABC"`, 1 * (1 + 2 + 3)},
		{`"Z", "ABC"`, 2*26 + 1*(1+2+3)},
	}

	for _, ex := range examples {
		result := Run(ex.input)
		if !reflect.DeepEqual(ex.answer, result) {
			t.Errorf("\nresult: %+v\nexpected: %d\ninput: %+v", result, ex.answer, ex.input)
		}
	}
}

// COLIN ... is the 938th name in the list
func TestParseSort(t *testing.T) {

	dat, err := ioutil.ReadFile("./names.txt")
	if err != nil {
		panic(err)
	}

	names := Parse(string(dat))
	if names[0] != "MARY" {
		t.Error("error parsing")
	}

	Sort(names)
	if names[937] != "COLIN" {
		t.Error("error sorting: " + names[937])
	}

}

func TestLetterScore(t *testing.T) {

	var examples = []struct {
		input  string
		answer int
		err    bool
	}{
		{"@", 0, true},
		{"!", 0, true},
		{"a", 1, false},
		{"A", 1, false},
		{"z", 26, false},
		{"Z", 26, false},
		{"C", 3, false},
		{"O", 15, false},
		{"L", 12, false},
		{"I", 9, false},
		{"N", 14, false},
	}

	for _, ex := range examples {
		result, err := LetterScore(ex.input[0])
		if ex.err && err == nil {
			t.Errorf("should be an error: %+v", ex.input)
		} else if !reflect.DeepEqual(ex.answer, result) {
			t.Errorf("\nresult: %+v\nexpected: %d\ninput: %+v", result, ex.answer, ex.input)
		}
	}
}
