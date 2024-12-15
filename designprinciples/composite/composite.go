package composite

import "fmt"

type IComponent interface {
	Draw()
}

type Circle struct {
	Name string
}

type Rectangle struct {
	Name string
}

func (c *Circle) Draw() {
	fmt.Printf("Drawing component :%s \n", c.Name)
}

func (r *Rectangle) Draw() {
	fmt.Printf("Drawing component :%s \n", r.Name)
}

type Composite struct {
	Name string
	Children []IComponent
}

func (c *Composite) Draw() {
	fmt.Printf("Drawing component :%s \n", c.Name)
	for _, child := range c.Children {
		child.Draw()
	}
}

func (c *Composite) AddChild(component IComponent) {
	c.Children = append(c.Children, component)
}

func (c *Composite) GetChildCount() int {
	return len(c.Children)
}
