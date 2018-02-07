// package main
/*

Using names.txt (right click and 'Save Link/Target As...'), a 46K text file containing over five-thousand first names, begin by sorting it into alphabetical order. Then working out the alphabetical value for each name, multiply this value by its alphabetical position in the list to obtain a name score.

For example, when the list is sorted into alphabetical order, COLIN, which is worth 3 + 15 + 12 + 9 + 14 = 53, is the 938th name in the list. So, COLIN would obtain a score of 938 × 53 = 49714.

What is the total of all the name scores in the file?

*/

package main

import (
	"fmt"
	"io/ioutil"
)

// Run run
func Run(input string) (answer int) {

	fmt.Println(input)

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
