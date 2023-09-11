package main

import (
	"fmt"
)

// struct
type Person struct {
	name          string
	age           uint
	languageSpoke []string
	familyMap     map[string]string
	clothing      Clothing
}

type Clothing struct {
	shoeSize   uint
	shirtColor string
}

// getter method for the above struct
func (p Person) GetName() string {
	return p.name
}

func main() {

	p1 := Person{}
	fmt.Println(p1)

	p2 := Person{
		"John", 34,
		[]string{"Eng, Rus"},
		map[string]string{
			"father": "Doe", "Mother": "Susan",
		},
		Clothing{
			10, "Blue",
		},
	}
	fmt.Println(p2)
	fmt.Println(p2.languageSpoke)

	name := p2.GetName()
	fmt.Println(name)

	fmt.Println("Accessing the embedded struct :", p2.clothing.shoeSize)
}
