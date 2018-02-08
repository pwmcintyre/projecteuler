// package main
/*

Using names.txt (right click and 'Save Link/Target As...'), a 46K text file containing over five-thousand first names, begin by sorting it into alphabetical order. Then working out the alphabetical value for each name, multiply this value by its alphabetical position in the list to obtain a name score.

For example, when the list is sorted into alphabetical order, COLIN, which is worth 3 + 15 + 12 + 9 + 14 = 53, is the 938th name in the list. So, COLIN would obtain a score of 938 × 53 = 49714.

What is the total of all the name scores in the file?

*/

package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

// Run run
func Run(input string) (answer int) {

	names := Parse(input)
	Sort(names)

	for i, name := range names {
		score := 0
		for i := range name {
			s, err := LetterScore(name[i])
			if err != nil {
				panic("found non-letter " + name)
			}
			score += s
		}
		score *= i + 1
		answer += score
	}

	return
}

const UPPER_MIN = 65
const LOWER_MIN = 97

func LetterScore(letter byte) (int, error) {
	if letter >= UPPER_MIN && letter < UPPER_MIN+26 {
		return int(letter - UPPER_MIN + 1), nil
	} else if letter >= LOWER_MIN && letter < LOWER_MIN+26 {
		return int(letter - LOWER_MIN + 1), nil
	}
	return 0, errors.New("not a letter: " + string(letter))
}

func Sort(slice []string) {
	sort.Strings(slice)
}

func Parse(input string) (names []string) {

	// clean
	input = strings.Replace(input, `"`, "", -1)
	input = strings.Replace(input, ` `, "", -1)

	// split
	names = strings.Split(input, ",")

	return
}

func main() {
	dat, err := ioutil.ReadFile("./problem22/names.txt")
	if err != nil {
		panic(err)
	}
	data := string(dat)
	fmt.Println(Run(data))
}
