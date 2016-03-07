package main

import "fmt"

type FlyBehavior interface {
	Fly()
}

type QuackBehavior interface {
	Quack()
}

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

func (t Duck) Swim() {
	fmt.Printf("All ducks can swim, even decoys!")
}

func main() {
	fmt.Printf("Strategy pattern in Go")
}
