package encapsulation

import "fmt"

var (
	CompanyName     = "Test"
	companyLocation = "somewhere"
)

type Person struct {
	Name string
	age  int
}

func (p *Person) getName() string {
	return p.Name
}

func (p *Person) GetAge() int {
	return p.age
}

type company struct {
}

func getCompanyName() string {
	return CompanyName
}

func GetCompanyLocation() string {
	return companyLocation
}

func GetPerson() *Person {
	p := &Person{
		Name: "test",
		age:  21,
	}
	fmt.Printf("Model Package : Person name : %s, age : %d", p.getName(), p.GetAge())
	return p
}
