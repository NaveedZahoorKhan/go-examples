package main

import "fmt"

func main() {
	a := [5]int{1, 1, 0, -1, -1}
	plusMinus(a)
}
func plusMinus(a [5]int) {
	length := len(a)
	positives := 0
	negatives := 0
	zeroes := 0
	for i := 0; i < length; i++ {
		if a[i] > 0 {
			positives++
		} else if a[i] < 0 {
			negatives++
		} else {
			zeroes++
		}
	}

	fmt.Println(float64(positives) / float64(length))
	fmt.Println(float64(negatives) / float64(length))
	fmt.Println(float64(zeroes) / float64(length))
}
