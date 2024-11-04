package model

import (
	"fmt"
	"go-program/goconcepts/encapsulation"
)

func Test() {
	p := encapsulation.Person{
		Name: "Test",
		//age : 23,   // not accessable
	}

	fmt.Println(p)

	//c := encapsulation.company{}   // not accessable
	fmt.Println(p.GetAge())
	//fmt.Println(p.getName())   // not accessable

	fmt.Println(p.Name)
	//fmt.Println(p.age)     // not accessable

	fmt.Println(encapsulation.CompanyName)
	//fmt.Println(encapsulation.companyLocation)  // not accessable

}
