package main

func main() {
	md := NewMallardDuck()
	md.PerformFly()
	md.PerformQuack()

	model := NewModelDuck()
	model.PerformFly()
	model.SetFlyBehavior(NewFlyRocketPowered())
	model.PerformFly()
}
