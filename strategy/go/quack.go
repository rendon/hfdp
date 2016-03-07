package main

import "fmt"

type QuackBehavior interface {
	Quack()
}

type Quack struct{}

func (t Quack) Quack() {
	fmt.Printf("Quack!\n")
}

func NewQuack() *Quack {
	return &Quack{}
}

type Squeak struct{}

func (t Squeak) Quack() {
	fmt.Printf("Squeak!")
}

func NewSqueak() *Squeak {
	return &Squeak{}
}
