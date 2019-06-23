package functions

// Circulations return all combinations of the characters of a given string
func Circulations(in string) []string {

	if len(in) == 0 {
		return []string{""}
	}

	str := []byte(in)
	buff := []byte(in)

	permutations := make([]string, len(in))
	for i := range permutations {

		// iterate over characters in the string
		for c := range str {

			// offset by 1 each permutation
			i := (c + i) % len(str)
			buff[i] = str[c]

		}

		permutations[i] = string(buff)

	}

	return permutations
}
