package main

import "fmt"

type Me struct {
	Name string
	Age  int
}

func (p Me) Infor() {
	fmt.Printf("tên %s tuổi %d\n", p.Name, p.Age)
}

func main() {

	me := Me{Name: "Tung", Age: 20}
	person.Infor()
}
