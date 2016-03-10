package main

import (
	"container/list"
	"fmt"

	"github.com/rendon/hfdp/observer/go/observer"
)

type DisplayElement interface {
	Display()
}

// Weather data
type WeatherData struct {
	Observers   *list.List
	Temperature float64
	Humidity    float64
	Pressure    float64
}

func NewWeatherData() WeatherData {
	return WeatherData{
		Observers: list.New(),
	}
}

func (t WeatherData) RegisterObserver(o observer.Observer) {
	t.Observers.PushBack(o)
}

func (t WeatherData) RemoveObserver(o observer.Observer) {
	for e := t.Observers.Front(); e != nil; e = e.Next() {
		if e.Value == o {
			t.Observers.Remove(e)
			return
		}
	}
}

func (t WeatherData) NotifyObservers() {
	for e := t.Observers.Front(); e != nil; e = e.Next() {
		o := e.Value.(observer.Observer)
		o.Update(t.Temperature, t.Humidity, t.Pressure)
	}
}

func (t WeatherData) measurementsChanged() {
	t.NotifyObservers()
}

func (t WeatherData) SetMeasurements(temperature, humidity, pressure float64) {
	t.Temperature = temperature
	t.Humidity = humidity
	t.Pressure = pressure
	t.measurementsChanged()
}

// Statistics display
type StatisticsDisplay struct {
	MaxTemp     float64
	MinTemp     float64
	TempSum     float64
	NumReadings float64
	WeatherData observer.Subject
}

func NewStatisticsDisplay(s observer.Subject) *StatisticsDisplay {
	sd := &StatisticsDisplay{
		MinTemp:     200,
		WeatherData: s,
	}
	sd.WeatherData.RegisterObserver(sd)
	return sd
}

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

func (t StatisticsDisplay) Display() {
	fmt.Print("Avg/Max/Min temperature = ")
	fmt.Printf("%f/%f/%f\n", t.TempSum/t.NumReadings, t.MaxTemp, t.MinTemp)
}

func main() {
	var wd observer.Subject = NewWeatherData()
	NewStatisticsDisplay(wd)
	wd.(WeatherData).SetMeasurements(80, 65, 30.4)
	wd.(WeatherData).SetMeasurements(82, 70, 29.2)
	wd.(WeatherData).SetMeasurements(78, 90, 29.2)
}
