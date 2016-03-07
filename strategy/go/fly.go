package main

import "fmt"

// FlyBehavior
type FlyBehavior interface {
	Fly()
}

// FlyNoWay
type FlyNoWay struct{}

func (t FlyNoWay) Fly() {
	fmt.Printf("I can't fly :(\n")
}

func NewFlyNoWay() *FlyNoWay {
	return &FlyNoWay{}
}

// FlyWithWings
type FlyWithWings struct{}

func (t FlyWithWings) Fly() {
	fmt.Printf("I'm flying!\n")
}

func NewFlyWithWings() *FlyWithWings {
	return &FlyWithWings{}
}

// FlyRocketPowered
type FlyRocketPowered struct{}

func (t FlyRocketPowered) Fly() {
	fmt.Printf("I'm flying with a rocket!\n")
}

func NewFlyRocketPowered() *FlyRocketPowered {
	return &FlyRocketPowered{}
}
