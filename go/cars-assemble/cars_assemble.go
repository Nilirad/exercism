// Package cars contains utility functions
// for calculating car manufacturing productivity.
package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * successRate / 100.0
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	return int(CalculateWorkingCarsPerHour(productionRate, successRate) / 60.0)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	const (
		bundleSize     uint = 10
		bundleCost     uint = 95_000
		individualCost uint = 10_000
	)

	bundles := uint(carsCount) / bundleSize
	individualCars := uint(carsCount) % bundleSize

	return bundles*bundleCost + individualCars*individualCost
}
