package main

import (
	"fmt"
	"math/rand"
)

//heap

type Heap struct {
	heap []int
}

// extract from heap
func (h *Heap) Extract() int {
	if len(h.heap) == 0 {
		return 0
	}
	max := h.heap[0]
	h.heap[0] = h.heap[len(h.heap)-1]
	h.heap = h.heap[:len(h.heap)-1]
	h.heapifyDown(0)
	return max
}

// insert into heap
func (h *Heap) Insert(val int) {
	h.heap = append(h.heap, val)
	h.heapifyUp(len(h.heap) - 1)
}

// heapify up

func (h *Heap) heapifyUp(i int) {
	for i > 0 {
		parent := (i - 1) / 2
		if h.heap[i] > h.heap[parent] {
			h.heap[i], h.heap[parent] = h.heap[parent], h.heap[i]
			i = parent
		} else {
			break
		}
	}
}

//heapify down

func (h *Heap) heapifyDown(i int) {
	for i < len(h.heap) {
		left := 2*i + 1
		right := 2*i + 2
		largest := i
		if left < len(h.heap) && h.heap[left] > h.heap[largest] {
			largest = left
		}
		if right < len(h.heap) && h.heap[right] > h.heap[largest] {
			largest = right
		}
		if largest != i {
			h.heap[i], h.heap[largest] = h.heap[largest], h.heap[i]
			i = largest
		} else {
			break
		}
	}
}

// heapify
func (h *Heap) heapify(i int) {
	for {
		left := 2*i + 1
		right := 2*i + 2
		largest := i
		if left < len(h.heap) && h.heap[left] > h.heap[largest] {
			largest = left
		}
		if right < len(h.heap) && h.heap[right] > h.heap[largest] {
			largest = right
		}
		if largest == i {
			break
		}
		h.heap[i], h.heap[largest] = h.heap[largest], h.heap[i]
		i = largest
	}
}

func main() {

	h := &Heap{}
	rand.Seed(1000)
	for i := 0; i < 10; i++ {
		h.Insert(rand.Intn(1000) + 1)
	}

	fmt.Println(h)
	v := h.Extract()

	fmt.Println(v)
	fmt.Println(h)
}
