package numbers

import (
	"reflect"
	"testing"
)

var examples = []struct {
	input  int
	answer string
}{
	{-1, "negative one"},
	{0, "zero"},
	{1, "one"},
	{2, "two"},
	{10, "ten"},
	{15, "fifteen"},
	{50, "fifty"},
	{150, "one hundred and fifty"},
	{599, "five hundred and ninety nine"},
	{1234, "one thousand two hundred and thirty four"},
	{9000000111222333444, "nine quintillion one hundred and eleven billion two hundred and twenty two million three hundred and thirty three thousand four hundred and forty four"},
}

func TestWordify(t *testing.T) {
	for _, ex := range examples {
		result := Wordify(ex.input)
		if !reflect.DeepEqual(ex.answer, result) {
			t.Errorf("\nresult: %+v\nexpected: %v\ninput: %+v", result, ex.answer, ex.input)
		}
	}
}

var hundreds = []struct {
	input  int
	answer string
}{
	{100, "one hundred"},
	{101, "one hundred and one"},
	{111, "one hundred and eleven"},
	{122, "one hundred and twenty two"},
	{555, "five hundred and fifty five"},
	{999, "nine hundred and ninety nine"},
}

func TestHundreds(t *testing.T) {
	for _, ex := range hundreds {
		result := Hundreds(ex.input)
		if !reflect.DeepEqual(ex.answer, result) {
			t.Errorf("\nresult: %+v\nexpected: %v\ninput: %+v", result, ex.answer, ex.input)
		}
	}
}
