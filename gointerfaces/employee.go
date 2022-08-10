package main

import "fmt"

type EmployeeInterface interface {
	PrintName(name string)
	PrintSalary(salary int)
}

type Employee struct {
	Name   string
	Salary int
}

func (emp Employee) PrintName(name string) {
	fmt.Println(name)
}

func main() {
	var emp1 Employee

	emp1.PrintName("John")
}
