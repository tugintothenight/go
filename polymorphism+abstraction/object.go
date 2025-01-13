package polymorphism

import (
	"fmt"
)

type identities interface {
	Identity()
}

type Person struct {
	name string
	age  int
}

type Laptop struct {
	owner Person
	name  string
	price int
}

func (p Person) Identity() {
	fmt.Printf("is Human\n")
}

func (l Laptop) Identity() {
	fmt.Printf("is laptop\n")
}
