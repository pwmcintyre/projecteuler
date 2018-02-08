package numbers

import (
	"math"
	"strings"
)

// https://en.wikipedia.org/wiki/List_of_numbers
// https://stackoverflow.com/questions/3299619/algorithm-that-converts-numeric-amount-into-english-words

var to_19 = []string{
	"zero",
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
	"ten",
	"eleven",
	"twelve",
	"thirteen",
	"fourteen",
	"fifteen",
	"sixteen",
	"seventeen",
	"eighteen",
	"nineteen",
}
var tens = []string{
	"",
	"",
	"twenty",
	"thirty",
	"forty",
	"fifty",
	"sixty",
	"seventy",
	"eighty",
	"ninety",
}
var denom = []string{
	"",
	"thousand",
	"million",
	"billion",
	"trillion",
	"quadrillion",
	"quintillion",
	// "sextillion", // need big numbers from here
	// "septillion",
	// "octillion",
	// "nonillion",
	// "decillion",
	// "undecillion",
	// "duodecillion",
	// "tredecillion",
	// "quattuordecillion",
	// "sexdecillion",
	// "septendecillion",
	// "octodecillion",
	// "novemdecillion",
	// "vigintillion",
}

func Wordify(n int) string {

	if n < 0 {
		return "negative " + Wordify(-n)
	}

	if n < 20 {
		return to_19[n]
	}

	words := ""
	for i := len(denom) - 1; i >= 0; i-- {
		magnitude := int(math.Pow(1000, float64(i)))
		if n >= magnitude {
			amount := int(n / magnitude)

			hundreds := Hundreds(amount)
			hundreds = hundreds + " " + denom[i]

			n = n - (amount * magnitude)
			words = words + " " + hundreds
		}
	}

	words = strings.TrimSpace(words)
	return words

}

func Hundreds(n int) (words string) {

	if n < 20 {
		words = words + to_19[n]
		return
	}

	if n >= 100 {
		hundreds := int(n / 100)
		words = to_19[hundreds] + " hundred"
		n = n - (hundreds * 100)
	}

	if n == 0 {
		return
	}
	if words != "" {
		words += " and "
	}
	if n < 20 {
		words = words + to_19[n]
		return
	}

	ten := int(n / 10)
	words = words + tens[ten]

	n = n - (ten * 10)
	if n == 0 {
		return
	}
	if n < 20 {
		words = words + " " + to_19[n]
		return
	}

	return
}
