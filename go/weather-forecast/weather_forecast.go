// Package weather provide some function to forcasting weather.
package weather


var (
	// CurrentCondition showing condition of the forcasting condition.
	CurrentCondition string
	// CurrentLocation showing location of the forcasting location.
	CurrentLocation  string
)

// Forecast return current forcast of specific city and condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
