package main

import "fmt"

func main() {
	// n := int32(2)
	// rows := int32(5)
	// lastAnswer := int32(0)
	// q := [5][3]int32{
	// {1, 0, 5},
	// {1, 1, 7},
	// {1, 0, 3},
	// {2, 1, 0},
	// {2, 1, 1},
	// }
	var arr = make([][]int32, 0)
	arr = append(arr, []int32{1, 0, 5})
	row := arr[0][:]
	row = append(row, 5)
	arr[0] = row
	fmt.Println(row)
}
