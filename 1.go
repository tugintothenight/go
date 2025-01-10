package main
import "fmt"

type Me struct {
	name string
	age int
}

func (m Me) Infor {
	fmt.Printf("My name is %s, I'm %d years old", m.name, m.age)
}

func main() {
	me := Me{"Tung", 20}
	me.Infor()
}