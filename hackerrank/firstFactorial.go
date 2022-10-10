package main

import "fmt"

func main() {
	num := 8
	f := firstFactorial(num)
	fmt.Println(f)
}

func firstFactorial(num int) int {
	// calculate the factorial of num
	if num == 0 {
		return 1
	}
	return num * firstFactorial(num-1)
}
