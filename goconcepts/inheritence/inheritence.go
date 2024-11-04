package inheritence

import "fmt"

type base struct {
	color string
}

func (b *base) say() {
	fmt.Printf("Hi from base color : %s\n", b.color)
}

type child struct {
	base
	style string
}

func check(b base) {
	b.say()
}

func TestStructInheritence() string {
	base := base{color: "red"}
	child := &child{
		base:  base,
		style: "Curve",
	}
	child.say() // we can call base function from child

	// check(child) // this will throw error
	return child.color
}
