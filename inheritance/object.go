package inheritance

import "fmt"

type Person struct {
	Name string
	Age  int
}

type Laptop struct {
	Owner Person
	Brand string
}

func NewLaptop(Name string, Age int, Brand string) Laptop {
	owner := Person{Name: Name, Age: Age}
	return Laptop{Owner: owner, Brand: Brand}
}
func NewPerson(Name string) Person {
	return Person{Name: Name}
}
func (l *Laptop) I() {
	fmt.Println("Laptop:", l.Owner, "Brand: ", l.Brand)
}

func (p *Person) I() {
	fmt.Printf("owner: name=%s\n", p.Name)
}
