package factory

import "fmt"

type IShape interface {
	Draw()
	GetName() string
}

type Circle struct {
	Name string
}

type Rectangle struct {
	Name string
}

type Square struct {
	Name string
}

func (d Circle) Draw() {
	fmt.Printf("Drawing : %s \n" , d.Name)
}
func (d Rectangle) Draw() {
	fmt.Printf("Drawing : %s \n" , d.Name)
}
func (d Square) Draw() {
	fmt.Printf("Drawing : %s \n" , d.Name)
}
func (d Circle) GetName() string {
	return d.Name
}
func (d Rectangle)GetName() string {
	return d.Name
}
func (d Square) GetName() string {
	return d.Name
}
type ShapeFactory struct {
}


func (s ShapeFactory) GetShape(input string) IShape {
	switch input {
	case "circle" :
		return Circle{
			Name: input,
		}
	case "rectangle" :
		return Rectangle{
			Name: input,
		}
	case "square" :
		return Square{
			Name: input,
		}
	}
	return Circle{
		Name: input,
	}
}