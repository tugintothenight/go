package main

import (
	"exp/encapsulation"
	"exp/inheritance"
	abstrpoly "exp/polymorphism+abstraction"
	"fmt"
)

func main() {
	// đóng gói
	me := encapsulation.NewPerson("Tung", 20)
	me.Grow()
	me.I()
	// đa hình + trừu tượng
	human := abstrpoly.Person{}
	dell := abstrpoly.Laptop{}
	human.Identity()
	dell.Identity()
	// kế thừa
	Dell := inheritance.NewLaptop("Tung", 20, "Dell")
	Dell.I()
	fmt.Println(Dell.Person.Name)
}
