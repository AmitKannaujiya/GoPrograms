package inheritence

import "fmt"

type iAnimal interface {
	breath()
}

type animal struct {
	name string
}

func (a *animal) breath() {
	fmt.Printf("Animal name : %s breath \n", a.name)
}

type iAquaticlAnimal interface {
	iAnimal
	swime()
}

type aquaticAnimal struct {
	animal
}

func (a *aquaticAnimal) swime() {
	fmt.Printf("Animal name : %s swime \n", a.name)
}

type iNonAquaticAnimal interface {
	iAnimal
	walk()
}
type nonAquaticAnimal struct {
	animal
}

func (a *nonAquaticAnimal) walk() {
	fmt.Printf("Animal name : %s walk \n", a.name)
}

type shark struct {
	aquaticAnimal
}

type lion struct {
	nonAquaticAnimal
}

func TestHeirarchicalInheritence() string {
	animal1 := animal{
		name: "shark",
	}
	shark := &shark{}
	shark.animal = animal1

	checkAquatic(shark)
	checkAnimal(shark)
	lion := &lion{}
	animal2 := animal{
		name: "lion",
	}
	lion.animal = animal2

	checkAnimal(lion)
	checkNonAquatic(lion)
	return lion.name
}

func checkAquatic(a iAquaticlAnimal) {
	a.swime()
}
func checkNonAquatic(a iNonAquaticAnimal) {
	a.walk()
}
func checkAnimal(a iAnimal) {
	a.breath()
}
