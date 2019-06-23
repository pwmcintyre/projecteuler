package functions

// Permutations return all combinations of the characters of a given string
func Permutations(in string) []string {
	permutations := []string{}
	permute(&permutations, in, 0)
	return permutations
}

// tips
// https://stackoverflow.com/questions/8306654/finding-all-possible-permutations-of-a-given-string-in-python
func permute(permutations *[]string, str string, step int) {

	// done
	if step >= len(str) {
		*permutations = append(*permutations, str)
		return
	}

	for i := step; i <= len(str)-1; i++ {

		chars := []byte(str)

		// swap
		t := chars[step]
		chars[step] = chars[i]
		chars[i] = t

		// recurse
		permute(permutations, string(chars), step+1)

	}

}
