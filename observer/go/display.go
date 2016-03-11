package main

import (
	"fmt"
	"math"

	"github.com/rendon/hfdp/observer/go/observer"
)

// DisplayElement display device interface
type DisplayElement interface {
	Display()
}

// StatisticsDisplay implements observer.Observer and DisplayElement.
type StatisticsDisplay struct {
	MaxTemp     float64
	MinTemp     float64
	TempSum     float64
	NumReadings float64
	WeatherData observer.Subject
}

// NewStatisticsDisplay returns a new StatisticsDisplay already registered
// to WeatherData.
func NewStatisticsDisplay(s observer.Subject) *StatisticsDisplay {
	sd := &StatisticsDisplay{
		MinTemp:     200,
		WeatherData: s,
	}
	sd.WeatherData.RegisterObserver(sd)
	return sd
}

// Update implements observer.Observer interface
func (t *StatisticsDisplay) Update(temperature, humidity, pressure float64) {
	t.TempSum += temperature
	if temperature > t.MaxTemp {
		t.MaxTemp = temperature
	}
	if temperature < t.MinTemp {
		t.MinTemp = temperature
	}
	t.Display()
}

// Display implements DisplayElement interface
func (t StatisticsDisplay) Display() {
	fmt.Print("Avg/Max/Min temperature = ")
	fmt.Printf("%f/%f/%f\n", t.TempSum/t.NumReadings, t.MaxTemp, t.MinTemp)
}

// CurrentConditionsDisplay a device to display the current conditions
type CurrentConditionsDisplay struct {
	Temperature float64
	Humidity    float64
	WeatherData observer.Subject
}

// NewCurrentConditionsDisplay returns a new CurrentConditionsDisplay already
// registered to WeatherData
func NewCurrentConditionsDisplay(s observer.Subject) *CurrentConditionsDisplay {
	d := &CurrentConditionsDisplay{
		WeatherData: s,
	}
	d.WeatherData.RegisterObserver(d)
	return d
}

// Update implements observer.Observer interface
func (t CurrentConditionsDisplay) Update(temp, h, p float64) {
	t.Temperature = temp
	t.Humidity = h
	t.Display()
}

// Display implements DisplayElement interface
func (t CurrentConditionsDisplay) Display() {
	fmt.Printf("Current conditions: %fF degrees and %f%% humidity\n",
		t.Temperature, t.Humidity)
}

// ForecastDisplay a device to display the weather forecast
type ForecastDisplay struct {
	CurrentPressure float64
	LastPressure    float64
	WeatherData     observer.Subject
}

// NewForecastDisplay returns a new ForecastDisplay already registered to
// WeatherData
func NewForecastDisplay(s observer.Subject) *ForecastDisplay {
	d := &ForecastDisplay{
		CurrentPressure: 29.92,
		WeatherData:     s,
	}
	d.WeatherData.RegisterObserver(d)
	return d
}

// Update implements observer.Observer interface
func (t *ForecastDisplay) Update(temp, h, p float64) {
	t.LastPressure = t.CurrentPressure
	t.CurrentPressure = p
	t.Display()
}

// Display implements DisplayElement interface
func (t ForecastDisplay) Display() {
	fmt.Printf("Forecast: ")
	if t.CurrentPressure > t.LastPressure {
		fmt.Println("Improving weather on the way")
	} else if math.Abs(t.CurrentPressure-t.LastPressure) < 1e-6 {
		fmt.Println("More of the same")
	} else if t.CurrentPressure < t.LastPressure {
		fmt.Println("Watch out for cooler, rainy weather!")
	}
}
