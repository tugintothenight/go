package main

import (
	"exp/encapsulation"
	"exp/inheritance"
	abstrpoly "exp/polymorphism+abstraction"
)

func main() {
	// đóng gói
	me := encapsulation.NewPerson("Tung", 20)
	me.Grow()
	me.I()
	// đa hình
	human := abstrpoly.Person{}
	dell := abstrpoly.Laptop{}
	human.Identity()
	dell.Identity()
	// kế thừa
	apple := inheritance.NewLaptop("Tung", 20, "Dell")
	apple.I()
	// Trừu tượng
}
