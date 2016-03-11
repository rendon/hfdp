package main

import (
	"container/list"

	"github.com/rendon/hfdp/observer/go/observer"
)

// WeatherData implements observer.Subject
type WeatherData struct {
	Observers   *list.List
	Temperature float64
	Humidity    float64
	Pressure    float64
}

// NewWeatherData returns a new WeatherData
func NewWeatherData() WeatherData {
	return WeatherData{
		Observers: list.New(),
	}
}

// RegisterObserver implements observer.Subject method
func (t WeatherData) RegisterObserver(o observer.Observer) {
	t.Observers.PushBack(o)
}

// RemoveObserver implements observer.Subject method
func (t WeatherData) RemoveObserver(o observer.Observer) {
	for e := t.Observers.Front(); e != nil; e = e.Next() {
		if e.Value == o {
			t.Observers.Remove(e)
			return
		}
	}
}

// NotifyObservers implements observer.Subject method
func (t WeatherData) NotifyObservers() {
	for e := t.Observers.Front(); e != nil; e = e.Next() {
		o := e.Value.(observer.Observer)
		o.Update(t.Temperature, t.Humidity, t.Pressure)
	}
}

func (t WeatherData) measurementsChanged() {
	t.NotifyObservers()
}

// SetMeasurements receives the the weathre measurements
func (t WeatherData) SetMeasurements(temperature, humidity, pressure float64) {
	t.Temperature = temperature
	t.Humidity = humidity
	t.Pressure = pressure
	t.measurementsChanged()
}
