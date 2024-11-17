package deccorator

import "fmt"

type BasePizza interface {
	Cost() int
	GetName() string
}

type FarmHouse struct {
	Name string
}

func (f *FarmHouse) Cost() int {
	return 100
}

func (f *FarmHouse) GetName() string {
	return f.Name
}

func NewFarmHouse(name string) *FarmHouse {
	return &FarmHouse{
		Name: name,
	}
}

type Margarita struct {
	Name string
}

func (m *Margarita) GetName() string {
	return m.Name
}

func NewMargrita(name string) *Margarita {
	return &Margarita{
		Name: name,
	}
}

func (f *Margarita) Cost() int {
	return 120
}

type Topping struct {
	Pizza BasePizza
}

type CheezeTopping struct {
	Topping Topping
}

type MashroomTopping struct {
	Topping Topping
}

func (c *CheezeTopping) Cost() int {
	return c.Topping.Pizza.Cost() + 10
}

func (m *MashroomTopping) Cost() int {
	return m.Topping.Pizza.Cost() + 20
}

func (c *CheezeTopping) ToString() string {
	return fmt.Sprintf("Pizza name %s, cost of pizza %d \n", c.Topping.Pizza.GetName(), c.Cost())
}

func (m *MashroomTopping) ToString() string {
	return fmt.Sprintf("Pizza name %s, cost of pizza %d \n", m.Topping.Pizza.GetName(), m.Cost())
}
