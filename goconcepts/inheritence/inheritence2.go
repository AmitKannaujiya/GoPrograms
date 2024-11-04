package inheritence

import "fmt"

// interface
type iBase interface {
	say2()
}

type base2 struct {
	color string
}

type child2 struct {
	base2
	style string
}

func (b *base2) say2() {
	fmt.Printf("Hi from color : %s\n", b.color)
	b.clear() // here both child and base both have clear base clear is called instead of child, opposite to java
}
func (b *base2) clear() {
	fmt.Println("Clear method of base is called")
}

func (b *child2) clear() {
	fmt.Println("Clear method of child is called")
}

func check2(base iBase) {
	base.say2()
}

func TestInterfaceInheritence() string {
	base := base2{color: "yellow"}
	child := &child2{
		base2: base,
		style: "paper work",
	}
	child.say2()
	check2(child)
	return child.color
}
