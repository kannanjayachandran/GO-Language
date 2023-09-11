package main

import "fmt"

type Name interface {
	GetName() string
}

type Person struct {
	firstName string
	lastName  string
}

func (p Person) GetName() string {
	return p.firstName + " " + p.lastName
}

func main() {

	var name Name = Person{"King", "Kong"}
	fmt.Println(name)
	fmt.Println(name.GetName())
}
