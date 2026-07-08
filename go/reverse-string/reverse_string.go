package reversestring

func Reverse(input string) string {
	runes := []rune(input)

	reversed := make([]rune, 0, len(runes))
	for i := len(runes) - 1; i >= 0; i-- {
		reversed = append(reversed, runes[i])
	}

	return string(reversed)
}
