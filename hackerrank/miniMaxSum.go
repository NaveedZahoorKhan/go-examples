package main

import "fmt"

func main() {
	a := [5]int32{256741038, 623958417, 467905213, 714532089, 938071625}
	s := []int32{}
	length := len(a)
	for j := 0; j < length; j++ {
		sum := int32(0)
		for i := 0; i < length; i++ {
			if i == j {
				continue
			}
			sum += a[i]
		}
		s = append(s, sum)
	}
	for j := 0; j < length-1; j++ {
		for i := 0; i < length-1; i++ {
			num1 := s[i]
			num2 := s[i+1]
			if num1 > num2 {
				s[i+1], s[i] = s[i], s[i+1]
			}
		}
	}
	fmt.Println(s)
	fmt.Printf("%d %d", s[0], s[length-1])
}
