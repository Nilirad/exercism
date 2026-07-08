package anagram

import (
	"slices"
	"strings"
	"unicode"
	"unicode/utf8"
)

// Detect returns the candidates that are anagrams of subject.
func Detect(subject string, candidates []string) []string {
	subjectRuneCount := utf8.RuneCountInString(subject)

	// Convert subject to lowercase runes and sort them
	subjectRunes := make([]rune, 0, subjectRuneCount)
	for _, r := range subject {
		subjectRunes = append(subjectRunes, unicode.ToLower(r))
	}
	slices.Sort(subjectRunes)

	// Pre-allocate a buffer for candidate runes to avoid allocations in the loop
	candidateRunesBuf := make([]rune, subjectRuneCount)

	var anagrams []string
	for _, candidate := range candidates {
		if strings.EqualFold(subject, candidate) {
			continue
		}

		if utf8.RuneCountInString(candidate) != subjectRuneCount {
			continue
		}

		candidateRunes := candidateRunesBuf[:0]
		for _, r := range candidate {
			candidateRunes = append(candidateRunes, unicode.ToLower(r))
		}

		slices.Sort(candidateRunes)

		if slices.Equal(subjectRunes, candidateRunes) {
			anagrams = append(anagrams, candidate)
		}
	}

	return anagrams
}
