package main

import "fmt"

func main() {
	n := 1
	for i := 6; i > 0; i-- {

		for j := 0; j < i-1; j++ {
			fmt.Print(" ")
		}
		for j := 0; j < n; j++ {
			fmt.Print("#")
		}
		fmt.Println("")
		n++
	}
}
