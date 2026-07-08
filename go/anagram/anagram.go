package anagram

import (
	"maps"
	"strings"
)

func getComposition(s string) map[rune]int {
	composition := make(map[rune]int)
	for _, r := range s {
		composition[r]++
	}

	return composition
}

func Detect(subject string, candidates []string) []string {
	subjectLower := strings.ToLower(subject)
	subjectComposition := getComposition(subjectLower)

	var anagrams []string
	for _, candidate := range candidates {
		candidateLower := strings.ToLower(candidate)
		if subjectLower == candidateLower {
			continue
		}

		candidateComposition := getComposition(candidateLower)
		if maps.Equal(subjectComposition, candidateComposition) {
			anagrams = append(anagrams, candidate)
		}
	}

	return anagrams
}
