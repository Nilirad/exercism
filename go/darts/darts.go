package darts

import "math"

const (
	smallCircle        = 1.0
	mediumCircle       = 5.0
	bigCircle          = 10.0
	smallCirclePoints  = 10
	mediumCirclePoints = 5
	bigCirclePoints    = 1
	missPoints         = 0
)

func Score(x, y float64) int {
	distance := math.Sqrt(x*x + y*y)

	if distance <= smallCircle {
		return smallCirclePoints
	}
	if distance <= mediumCircle {
		return mediumCirclePoints
	}
	if distance <= bigCircle {
		return bigCirclePoints
	}
	return missPoints
}
