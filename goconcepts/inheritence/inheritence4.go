package inheritence

import "fmt"

// interface
type iBase2 interface {
	say4()
}
type iBase3 interface {
	walk()
}
type base4 struct {
	color  string
	clear1 func()
}

func (b *base4) walk() {
	fmt.Printf("walk for color : %s\n", b.color)
	//b.clear1()   // here both child and base both have clear base clear is called instead of child, opposite to java
}

type base5 struct{}
type child4 struct {
	base4
	base5
	style string
}

func (b *base4) say4() {
	fmt.Printf("Hi from color : %s\n", b.color)
	b.clear1() // here both child and base both have clear base clear is called instead of child, opposite to java
}

// func(b *base3) clear1() {
// 	fmt.Println("Clear method of base is called")
// }

func (b *child4) clear1() {
	fmt.Println("Clear method of child is called")
}

func check4(base iBase2) {
	base.say4()
}

func check5(base iBase3) {
	base.walk()
}

func TestInterfaceInheritence3() string {
	base := base4{color: "blue",
		clear1: func() {
			fmt.Println("Clear method of base is called")
		},
	}
	base2 := base5{}
	child := &child4{
		base4: base,
		base5: base2,
		style: "paper work",
	}
	child.say4()
	child.walk()
	check4(child)
	check5(child)
	return child.color
}
