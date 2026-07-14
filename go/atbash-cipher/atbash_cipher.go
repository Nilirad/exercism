package atbashcipher

import "strings"

const chunkSize = 5

func Atbash(s string) string {
	var builder strings.Builder
	builder.Grow(len(s) + (len(s)+chunkSize-1)/chunkSize + 1)

	counter := 0
	for i := range len(s) {
		char := s[i]
		if char >= 'A' && char <= 'Z' {
			char += 'a' - 'A'
		}

		switch {
		case char >= 'a' && char <= 'z':
			char = 'a' + 'z' - char
		case char >= '0' && char <= '9':
		default:
			continue
		}

		if counter >= chunkSize {
			builder.WriteByte(' ')
			counter = 0
		}
		builder.WriteByte(char)
		counter++
	}

	return builder.String()
}
