package main

import (
	"fmt"
	"sync"
)

var sharedVal = 1000

func main() {
	// By default go cares only for main thread, once it stops it will exit
	// So you need to create a wait group to wait for all goroutines to finish
	wg := sync.WaitGroup{}

	// Create a mutex to lock sharedVal
	m := sync.Mutex{}
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go reductor(&wg, &m)
	}
	wg.Wait()
	fmt.Println(sharedVal)
}
func reductor(wg *sync.WaitGroup, m *sync.Mutex) {
	defer wg.Done()
	m.Lock()
	sharedVal--

	m.Unlock()
}
