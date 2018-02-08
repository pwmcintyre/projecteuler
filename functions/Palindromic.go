package functions

import (
	"strconv"
)

func IsPalindromic(input int) bool {
	str := strconv.Itoa(input)
	for i, _ := range str {
		if str[i] != str[len(str)-1-i] {
			return false
		}
	}
	return true
}
