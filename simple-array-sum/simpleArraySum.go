package main

import "fmt"

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	s := getSum(0, a[0], a)
	fmt.Println(s)
}
func getSum(index, sum int, a [5]int) int {
	index += 1
	if index >= len(a) {
		return sum
	}
	return getSum(index, a[index]+sum, a)
}
