package main

import "fmt"

// person store has changeable Name field via export type
type person struct {
	Name string
}

type person2 struct {
	name string
}

func (p *person2) Name() string {
	return p.name
}

func (p *person2) SetName(name string) {
	p.name = name
}

func main() {
	p := &person{Name: "Roman"}
	p.Name = "Olga"
	fmt.Println(p.Name)

	p2 := person2{name: "Roman"}
	fmt.Println(p2.Name())
	p2.SetName("Lesha")
	fmt.Println(p2.Name())
}
