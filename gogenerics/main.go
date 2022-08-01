package main

import "fmt"

func main() {

	testAge := 23
	testAge2 := 25
	add[int](testAge, testAge2)
}

func add[T int | float64](a T, b T) {
	if a < b {
		fmt.Println(a + b)
	} else {
		fmt.Println(b + a)
	}
}
