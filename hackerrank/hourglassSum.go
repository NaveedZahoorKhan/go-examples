package main

import "fmt"

func main() {
	arr := [][]int32{
		{-1, -1, 0, -9, -2, -2},
		{-2, -1, -6, -8, -2, -5},
		{-1, -1, -1, -2, -3, -4},
		{-1, -9, -2, -4, -4, -5},
		{-7, -3, -3, -2, -9, -9},
		{-1, -3, -1, -2, -4, -5},
	}
	length := len(arr)
	const MaxUint = ^uint(0)
	const MaxInt = int(MaxUint >> 1)
	const MinInt = -MaxInt - 1

	var max int = MinInt
	for row := 0; row < length-2; row++ {
		for col := 0; col < length-2; col++ {
			s := arr[row][col] + arr[row][col+1] + arr[row][col+2] + arr[row+1][col+1] + arr[row+2][col] + arr[row+2][col+1] + arr[row+2][col+2]
			if int(s) > max {
				max = int(s)
			}
			fmt.Print(s)
			fmt.Print(" ")
		}
		fmt.Println("")
	}
	fmt.Println(max)

}
