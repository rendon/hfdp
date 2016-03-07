package main

import (
	"fmt"
	"os"
)

// Duck
type Duck struct {
	flyBehavior   FlyBehavior
	quackBehavior QuackBehavior
}

func (t Duck) PerformFly() {
	t.flyBehavior.Fly()
}

func (t Duck) PerformQuack() {
	t.quackBehavior.Quack()
}

func (t *Duck) SetFlyBehavior(fb FlyBehavior) {
	t.flyBehavior = fb
}

func (t Duck) Swim() {
	fmt.Printf("All ducks can swim, even decoys!\n")
}

func (t Duck) Display() {
	fmt.Fprintf(os.Stderr, "========== Not implemented ==========\n")
	os.Exit(1)
}

// ModelDuck
type ModelDuck struct {
	Duck
}

func (t ModelDuck) Display() {
	fmt.Printf("I'm a model duck!\n")
}

func NewModelDuck() *ModelDuck {
	return &ModelDuck{
		Duck{
			flyBehavior:   NewFlyNoWay(),
			quackBehavior: NewQuack(),
		},
	}
}

// MallardDuck
type MallardDuck struct {
	Duck
}

func (t MallardDuck) Display() {
	fmt.Printf("I'm a Mallard Duck!\n")
}

func NewMallardDuck() *MallardDuck {
	return &MallardDuck{
		Duck{
			quackBehavior: NewQuack(),
			flyBehavior:   NewFlyWithWings(),
		},
	}
}
