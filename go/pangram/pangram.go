package pangram

func IsPangram(input string) bool {
	const lettersInAlphabet = 26
	if len(input) < lettersInAlphabet {
		return false
	}

	var seen [lettersInAlphabet]bool

	// We can iterate over bytes instead of runes
	// because we are searching English letters which are all ASCII characters.
	for i := range len(input) {

		char := input[i]
		if char >= 'A' && char <= 'Z' {
			char += 'a' - 'A'
		}
		if char < 'a' || char > 'z' {
			continue
		}

		index := int(char - 'a')
		seen[index] = true
	}

	for i := range lettersInAlphabet {
		if !seen[i] {
			return false
		}
	}

	return true
}
