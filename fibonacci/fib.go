package main

import "fmt"

func main() {
	f := fib(10)
	fmt.Println("fib is", f)
	printFibSeries(10)
}
func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

func printFibSeries(n int) {
	a := 0
	b := 1
	c := b
	sum := b
	fmt.Printf("Series is: %d %d", a, b)
	for true {
		c = b
		b = a + b
		if b > n {
			fmt.Println()
			break
		}

		a = c
		sum += b
		fmt.Printf(" %d", b)

	}
	fmt.Println("Sum is:", sum)
}
