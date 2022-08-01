package main

import "fmt"

func reverseArray(a []int32) []int32 {
	// Write your code here
	for i := 0; i < len(a)/2; i++ {
		a[i], a[len(a)-1-i] = a[len(a)-1-i], a[i]
		fmt.Println(a[i], a[len(a)-i-1])
		fmt.Println(a)
	}
	return a
}

func main() {
	a := [9]int32{6676, 3216, 4063, 8373, 423, 586, 8850, 6762, 2}
	reverseArray(a[:])
}
