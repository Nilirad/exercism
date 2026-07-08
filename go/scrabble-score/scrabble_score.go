package scrabblescore

func Score(word string) int {
	score := 0
	for i := range len(word) {
		char := word[i]

		if char >= 'A' && char <= 'Z' {
			char += 'a' - 'A'
		}

		switch char {
		case 'a', 'e', 'i', 'o', 'u', 'l', 'n', 'r', 's', 't':
			score += 1
		case 'd', 'g':
			score += 2
		case 'b', 'c', 'm', 'p':
			score += 3
		case 'f', 'h', 'v', 'w', 'y':
			score += 4
		case 'k':
			score += 5
		case 'j', 'x':
			score += 8
		case 'q', 'z':
			score += 10
		}
	}
	return score
}
