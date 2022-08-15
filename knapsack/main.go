package main

func knapsack(items []int, capacity int) int {
	if len(items) == 0 {
		return 0
	}
	if capacity == 0 {
		return 0
	}
	if items[0] > capacity {
		return knapsack(items[1:], capacity)
	}
	return max(knapsack(items[1:], capacity), knapsack(items[1:], capacity-items[0])+items[0])
}
func max()
func main() {

}
