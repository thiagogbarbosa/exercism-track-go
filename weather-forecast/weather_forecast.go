//Package weather provides information about
//the current weather condition given a location.
package weather


//CurrentCondition storage the current weather condition.
var CurrentCondition string

//CurrentLocation storage the current location.
var CurrentLocation string

//Forecast provides the current weather condition in a giver location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
