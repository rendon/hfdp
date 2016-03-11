package main

import (
	"fmt"

	"github.com/rendon/hfdp/observer/go/observer"
)

func main() {
	var wd observer.Subject = NewWeatherData()
	sd := NewStatisticsDisplay(wd)
	ccd := NewCurrentConditionsDisplay(wd)
	NewForecastDisplay(wd)
	wd.(WeatherData).SetMeasurements(80, 65, 30.4)
	fmt.Println()
	wd.(WeatherData).SetMeasurements(82, 70, 29.2)
	fmt.Println()
	wd.(WeatherData).SetMeasurements(78, 90, 29.2)

	fmt.Println("\n===========================================================")
	wd.RemoveObserver(sd)
	wd.RemoveObserver(ccd)

	wd.(WeatherData).SetMeasurements(80, 65, 30.4)
	fmt.Println()
	wd.(WeatherData).SetMeasurements(82, 70, 29.2)
	fmt.Println()
	wd.(WeatherData).SetMeasurements(78, 90, 29.2)

}
