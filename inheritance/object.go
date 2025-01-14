package inheritance

import "fmt"

type Person struct {
	Name string
	Age  int
}

type Laptop struct {
	Person
	Brand string
}

func NewLaptop(Name string, Age int, Brand string) Laptop {
	owner := Person{Name: Name, Age: Age}
	return Laptop{Person: owner, Brand: Brand}
}
func NewPerson(Name string) Person {
	return Person{Name: Name}
}
func (l *Laptop) I() {
	fmt.Println("Laptop:", l.Person, "Brand: ", l.Brand)
}

func (p *Person) I() {
	fmt.Printf("owner: name=%s\n", p.Name)
}
