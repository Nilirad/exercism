package isogram

func IsIsogram(word string) bool {
	var seen [26]bool

	for i := range len(word) {
		charA := word[i]
		index, isAlpha := validateAndNormalize(charA)
		if !isAlpha {
			continue
		}

		if seen[index] {
			return false
		}
		seen[index] = true
	}

	return true
}

// validateAndNormalize ensures the byte is an ASCII letter
// and converts it to an alphabet index.
func validateAndNormalize(c byte) (int, bool) {
	if c >= 'A' && c <= 'Z' {
		c += 'a' - 'A'
	}
	if c < 'a' || c > 'z' {
		return 0, false
	}

	return int(c - 'a'), true
}
