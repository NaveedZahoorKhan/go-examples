package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p *Person) setPerson(name string, age int) {
	p.Age = age
	p.Name = name
}

func main() {
	p := Person{}
	p.setPerson("Naveed", 30)
	fmt.Println(p)
}
