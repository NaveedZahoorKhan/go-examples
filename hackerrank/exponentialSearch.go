package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []int32{1, 2, 4, 5, 8, 10, 12, 15, 19}
	search(arr, 10)
}

func search(arr []int32, v int32) {

	if arr[0] == v {
		fmt.Println(arr[0])
		return
	}
	n := len(arr)
	i := 1
	for i = 1; i < n && arr[i] < v; {
		i = i * 2
	}
	m := math.Min(float64(i), float64(n-1))
	x := bSearch(arr, int32(i/2), int32(m), v)

	fmt.Println(x)
}
func bSearch(arr []int32, low, high, x int32) int32 {
	if high > low {
		mid := low + (high-1)/2

		if arr[mid] == x {
			fmt.Println(x)
			return x
		}

		if arr[mid] > x {
			return bSearch(arr, low, mid-1, x)
		}

		return bSearch(arr, low, mid+1, x)
	}
	return -1

}
