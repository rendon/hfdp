package main

type Beverage interface {
	GetDescription() string
	Cost() float64
}

type CondimentDecorator struct {
	beverage Beverage
}

func (t CondimentDecorator) GetDescription() string {
	return "Unknown Beverage"
}

func (t CondimentDecorator) Cost() float64 {
	return 0
}

// Espresso
type Espresso struct{}

func (t Espresso) GetDescription() string {
	return "Espresso"
}

func (t Espresso) Cost() float64 {
	return 1.99
}

func NewEspresso() Beverage {
	return Espresso{}
}

// DarkRoast
type DarkRoast struct{}

func (t DarkRoast) GetDescription() string {
	return "Dark Roast"
}

func (t DarkRoast) Cost() float64 {
	return 0.99
}

func NewDarkRoast() Beverage {
	return DarkRoast{}
}

// HouseBlend
type HouseBlend struct{}

func (t HouseBlend) GetDescription() string {
	return "House Blend"
}

func (t HouseBlend) Cost() float64 {
	return 0.89
}

func NewHouseBlend() Beverage {
	return HouseBlend{}
}

// Mocha
type Mocha struct {
	*CondimentDecorator
}

func (t Mocha) GetDescription() string {
	return t.beverage.GetDescription() + ", Mocha"
}

func (t Mocha) Cost() float64 {
	return 0.20 + t.beverage.Cost()
}

func NewMocha(b Beverage) Beverage {
	return Mocha{&CondimentDecorator{beverage: b}}
}

// Whip
type Whip struct {
	*CondimentDecorator
}

func (t Whip) GetDescription() string {
	return t.beverage.GetDescription() + ", Whip"
}

func (t Whip) Cost() float64 {
	return 0.10 + t.beverage.Cost()
}

func NewWhip(b Beverage) Beverage {
	return Whip{&CondimentDecorator{beverage: b}}
}

// Soy
type Soy struct {
	*CondimentDecorator
}

func (t Soy) GetDescription() string {
	return t.beverage.GetDescription() + ", Soy"
}

func (t Soy) Cost() float64 {
	return 0.15 + t.beverage.Cost()
}

func NewSoy(b Beverage) Beverage {
	return Soy{&CondimentDecorator{beverage: b}}
}
