package highscores

import "math"

type HighScores struct {
	scores []int
}

// NewHighScores returns a new HighScores object.
func NewHighScores(scores []int) *HighScores {
	return &HighScores{scores: scores}
}

// Scores returns all the scores.
func (s *HighScores) Scores() []int {
	return s.scores
}

// Latest returns the latest (last) score.
func (s *HighScores) Latest() int {
	if len(s.scores) < 1 {
		return 0
	}

	return s.scores[len(s.scores)-1]
}

// PersonalBest returns the best (highest) score.
func (s *HighScores) PersonalBest() int {
	highest := math.MinInt

	for _, score := range s.scores {
		highest = max(highest, score)
	}

	return highest
}

// TopThree returns the top three scores.
func (s *HighScores) TopThree() []int {
	picked := make([]bool, len(s.scores))
	topThree := make([]int, 0, 3)

	for range 3 {
		pickedIndex := -1
		highest := math.MinInt
		for i, score := range s.scores {
			if picked[i] {
				continue
			}

			if score >= highest {
				highest = score
				pickedIndex = i
			}
		}

		if pickedIndex == -1 {
			break
		}

		picked[pickedIndex] = true
		topThree = append(topThree, highest)
	}

	return topThree
}
