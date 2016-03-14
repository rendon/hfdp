package main

import (
	"container/list"
	"fmt"
)

// region PizzaStore
// PizzaStore
type PizzaStore interface {
	CreatePizza(string) Pizza
	OrderPizza(string) Pizza
}

type PizzaStoreBase struct {
	v PizzaStore
}

func (t PizzaStoreBase) CreatePizza(typ string) Pizza {
	return nil
}

func (t PizzaStoreBase) OrderPizza(typ string) Pizza {
	pizza := t.v.CreatePizza(typ)
	fmt.Printf("--- Making a " + pizza.GetName() + " ---\n")
	pizza.Prepare()
	pizza.Bake()
	pizza.Cut()
	pizza.Box()
	return pizza
}

// endregion

// region Pizza
// Pizza
type Pizza interface {
	Prepare()
	Bake()
	Cut()
	Box()
	GetName() string
}

type PizzaBase struct {
	name     string
	dough    string
	sauce    string
	toppings *list.List
}

func (t PizzaBase) Prepare() {
	fmt.Printf("Prepare %s\n", t.name)
	fmt.Printf("Tossing dough...\n")
	fmt.Printf("Adding sauce...\n")
	fmt.Printf("Adding toppings:\n")
	for e := t.toppings.Front(); e != nil; e = e.Next() {
		fmt.Printf("    %s\n", e.Value.(string))
	}
}

func (t PizzaBase) Bake() {
	fmt.Printf("Bake for 25 minutes at 350\n")
}
func (t PizzaBase) Cut() {
	fmt.Printf("Cut the pizza into diagonal slices\n")
}

func (t PizzaBase) Box() {
	fmt.Printf("Place pizza in official PizzaStore box\n")
}

func (t PizzaBase) GetName() string {
	return t.name
}

// endregion

//region NYStyleCheesePizza
// NYStyleCheesePizza
type NYStyleCheesePizza struct {
	*PizzaBase
}

func NewNYStyleCheesePizza() Pizza {
	toppings := list.New()
	toppings.PushBack("Grated Reggiano Cheese")
	return NYStyleCheesePizza{
		&PizzaBase{
			name:     "NY Style Sauce and Cheese Pizza",
			dough:    "Thin Crust Dough",
			sauce:    "Marinara Sauce",
			toppings: toppings,
		},
	}
}

//endregion

// region NYPizzaStore
// NYPizzaStore
type NYPizzaStore struct {
	*PizzaStoreBase
}

func (t NYPizzaStore) CreatePizza(typ string) Pizza {
	if typ == "cheese" {
		return NewNYStyleCheesePizza()
	}
	return nil
}

func NewNYPizzaStore() PizzaStore {
	base := &PizzaStoreBase{}
	store := NYPizzaStore{base}
	store.v = store
	return store
}

// endregion

func main() {
	ny := NewNYPizzaStore()
	pizza := ny.OrderPizza("cheese")
	fmt.Printf("Ethan ordered %s\n\n", pizza.GetName())

	//ch := NewChicagoPizzaStore()
	//pizza = ch.OrderPizza("cheese")
	//fmt.Printf("Joel ordered %s\n\n", pizza.GetName())
}
