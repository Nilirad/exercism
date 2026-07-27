package etl

func Transform(in map[int][]string) map[string]int {
	out := make(map[string]int, 26)

	for score, letters := range in {
		for _, letter := range letters {
			out[string(letter[0]+'a'-'A')] = score
		}
	}

	return out
}
