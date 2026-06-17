// Package weather provides tools for weather forecasting.
package weather

var (
	// CurrentCondition represents the current weather condition.
	CurrentCondition string

	// CurrentLocation represents the current location.
	CurrentLocation string
)

// Forecast shows the weather forecast for a city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
