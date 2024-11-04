package inheritence

import "fmt"

// interface
type iBase1 interface {
	say3()
}

type base3 struct {
	color  string
	clear1 func()
}

type child3 struct {
	base3
	style string
}

func (b *base3) say3() {
	fmt.Printf("Hi from color : %s\n", b.color)
	b.clear1() // here both child and base both have clear base clear is called instead of child, opposite to java
}

// func(b *base3) clear1() {
// 	fmt.Println("Clear method of base is called")
// }

func (b *child3) clear1() {
	fmt.Println("Clear method of child is called")
}

func check3(base iBase1) {
	base.say3()
}

func TestInterfaceInheritence2() string {
	base := base3{color: "green",
		clear1: func() {
			fmt.Println("Clear method of base is called")
		},
	}
	child := &child3{
		base3: base,
		style: "paper work",
	}
	//child.say3()
	check3(child)
	return child.color
}
