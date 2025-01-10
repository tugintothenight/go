package main

import "fmt"

type Informer interface {
	Infor()
}

type Me struct {
	Name string
	Age  int
}

type Laptop struct {
	Owner Me
	Name  string
	Price int
}

func newMe(name string, age int) Me {
	return Me{Name: name, Age: age}
}

func newLaptop(nameOwner string, ageOwner int, nameLaptop string, price int) Laptop {
	me := newMe(nameOwner, ageOwner)
	laptop := Laptop{Owner: me, Name: nameLaptop, Price: price}
	return laptop
}

func (p Me) Infor() {
	fmt.Printf("name: %s age: %d\n", p.Name, p.Age)
}

func (l Laptop) Infor() {
	fmt.Printf("Owner: %s, Age: %d, Laptop: %s, Price: %d\n", l.Owner.Name, l.Owner.Age, l.Name, l.Price)
}

func printInfo(i Informer) {
	i.Infor()
}

func main() {
	me := newMe("Tung", 20)
	laptop := newLaptop("Tung", 21, "Dell", 1000)
	printInfo(me)
	printInfo(laptop)
}
