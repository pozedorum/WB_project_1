package main

import "fmt"

// текущая реализация на сервере
type FahrenheitThermometer struct{}

func (f *FahrenheitThermometer) GetFahrenheitTemp() float64 {
	return 77.0
}

// реализация требуемая потребителем
type CelsiusTemperature interface {
	GetCelsius() float64
}

func DisplayTemperature(temp CelsiusTemperature) {
	fmt.Printf("Текущая температура: %.1f°C\n", temp.GetCelsius())
}

// мой адаптер

type AdaptedThermometer struct {
	*FahrenheitThermometer
}

func (at *AdaptedThermometer) GetCelsius() float64 {
	return (at.GetFahrenheitTemp() - 32) * 5 / 9
}

func NewAdaptedThermometer(f *FahrenheitThermometer) *AdaptedThermometer {
	return &AdaptedThermometer{f}
}

// проверка
func main() {
	fahrenheitSensor := &FahrenheitThermometer{}
	adapter := NewAdaptedThermometer(fahrenheitSensor)
	DisplayTemperature(adapter)
}
