package isogram

func IsIsogram(word string) bool {
	var seen [26]bool

	for i := range len(word) {
		char := word[i]
		if char >= 'A' && char <= 'Z' {
			char += 'a' - 'A'
		}
		if char < 'a' || char > 'z' {
			continue
		}

		index := int(char - 'a')
		if seen[index] {
			return false
		}
		seen[index] = true
	}

	return true
}
