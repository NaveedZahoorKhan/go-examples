package main

import (
	"fmt"
	"math/rand"
)

func main() {
	rand.Seed(100000)
	candles := []int64{18, 90, 90, 13, 90, 75, 90, 8, 90, 43}
	length := len(candles)
	var max int64 = 0
	var maxCount int64 = 0
	var electedAt int = 0
	for j := 0; j < length; j++ {
		n := candles[j]
		if n > max {
			electedAt = j
			// new max is being elected
			max = candles[j]
			// reset count
			maxCount = 1
		}
		if n == max && electedAt != j {
			maxCount++
		}
	}

	fmt.Println(candles, max, maxCount)

}
