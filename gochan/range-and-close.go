package main

import "fmt"

func main() {
	c := make(chan int, 20)
	go fib(cap(c), c)

	for i := range c {
		fmt.Println(i)
	}
}

func fib(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}
