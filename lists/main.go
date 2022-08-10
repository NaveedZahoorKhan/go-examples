package main

import "fmt"

type Node struct {
	prev  *Node
	next  *Node
	value any
}
type List struct {
	head *Node
	tail *Node
	len  int
}

func (l *List) Init() {
	l.head = nil
	l.tail = nil
	l.len = 0
}
func (l *List) AddPrev(value any) {
	l.len++
	node := &Node{prev: nil, next: l.head, value: value}
	if l.tail == nil {
		l.tail = node
	} else {
		l.tail.prev = node
	}
}
func (l *List) AddNext(value any) {
	l.len++
	node := &Node{prev: l.tail, next: nil, value: value}
	if l.head == nil {
		l.head = node
	} else {
		l.head.next = node
	}
}
func (l *List) Remove() {
	l.len--
	node := l.Search("xyz")
	if node != nil {
		fmt.Println("Found")
	}
}
func (l *List) Search(key any) *Node {
	node := l.head
	for node != nil {
		if node.value == key {
			return node
		}
		node = node.next
	}
	return nil
}
func main() {
	l := &List{}
	l.Init()
	l.AddNext(5)
	l.AddPrev(100)
	l.AddNext(200)
	node := l.Search(100)

	if node != nil {
		fmt.Println("Found")
	}
}
