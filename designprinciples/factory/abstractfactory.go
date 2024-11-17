package factory

import "fmt"

type Vehicle interface {
	Drive()
	GetName() string
}

type VehicleFactory interface {
	GetVehicle() Vehicle
}

type LuxuryVehicleFactory struct{}

func (l *LuxuryVehicleFactory) GetVehicle(input string) Vehicle {
	switch input {
	case "bmw":
		return &BMW{
			Name: input,
		}
	case "merchadise":
		return &Mercharise{
			Name: input,
		}
	}
	return &BMW{
		Name: input,
	}
}

type OrdinaryVehicleFactory struct{
}

func (o *OrdinaryVehicleFactory) GetVehicle(input string) Vehicle {
	switch input {
	case "hundai":
		return &Hundai{
			Name: input,
		}
	case "swift":
		return &Swift{
			Name: input,
		}
	}
	return &Swift{
		Name: input,
	}
}

type BMW struct {
	Name string
}

func (b *BMW) Drive() {
	fmt.Printf("Driving : %s \n", b.Name)
}
func (b *BMW) GetName() string {
	return b.Name
}

type Mercharise struct {
	Name string
}

func (m *Mercharise) Drive() {
	fmt.Printf("Driving : %s \n", m.Name)
}

func (m *Mercharise) GetName() string {
	return m.Name
}

type Swift struct {
	Name string
}

func (s *Swift) Drive() {
	fmt.Printf("Driving : %s \n", s.Name)
}

func (s *Swift) GetName() string {
	return s.Name
}

type Hundai struct {
	Name string
}

func (h *Hundai) Drive() {
	fmt.Printf("Driving : %s \n", h.Name)
}

func (h *Hundai) GetName() string {
	return h.Name
}