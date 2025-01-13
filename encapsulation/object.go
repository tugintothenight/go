package encapsulation

import "fmt"

type Person struct {
	name string
	age  int
}

type Laptop struct {
	owner Person
	name  string
	price int
}

func NewPerson(name string, age int) Person {
	return Person{name: name, age: age}
}

func (p *Person) Grow() {
	p.age++
}

func NewLaptop(nameOwner string, ageOwner int, nameLaptop string, price int) Laptop {
	owner := NewPerson(nameOwner, ageOwner)
	return Laptop{owner: owner, name: nameLaptop, price: price}
}

func (p *Person) I() {
	fmt.Printf("Person: name=%s, age=%d\n", p.name, p.age)
}
