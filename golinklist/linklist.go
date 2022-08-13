package main

import "fmt"

// link list have nodes with value and next pointer
// link lists are in a ways similar to arrays but the are not adjacent indexes
// cost of inserting and deleting is O(n)
// cost of searching is O(n)

// adding a node on head has lower cost than arrays

// define a node

type node struct {
	value int
	next  *node
}

// define link list

type linkList struct {
	head   *node
	length int
}

// insert node on head
func (l *linkList) insertOnFront(v int) {
	n := &node{value: v}
	// link list is empty
	if l.length == 0 {
		l.head = n
		l.length++
	} else {
		// link list already have head
		n.next = l.head
		l.head = n
		l.length++
	}
}

// insert node on tail

func (l *linkList) insertOnTail(v int) {
	n := &node{value: v}

	if l.length == 0 {
		l.head = n
	} else {
		prev := l.head
		for i := 0; i < l.length; i++ {
			if prev.next == nil {
				prev.next = n
				l.length++
				return
			}
			prev = prev.next
		}
	}
}

// delete node
func (l *linkList) deleteNode(v int) {
	// want to delete head
	if l.head.value == v {
		// l is the only node
		if l.head.next == nil {
			l.head = nil
		} else {
			// l has more nodes
			l.head = l.head.next
			l.length--
		}

		return
	} else {
		// want to delete in between or tail
		prevNode := l.head
		for i := 0; i < l.length; i++ {
			if prevNode.next == nil {
				return
			}
			if prevNode.next.value == v {
				prevNode.next = prevNode.next.next
				l.length--
			}
			prevNode = prevNode.next
		}
	}

}

// search node

func (l *linkList) searchNode(v int) bool {

	prevNode := l.head

	// if value matches head
	if prevNode.value == v {
		return true
	}

	for i := 0; i < l.length; i++ {
		if prevNode.next == nil {
			return false
		}
		if prevNode.next.value == v {
			return true
		}

		prevNode = prevNode.next

	}
	return false
}

// print list

func (l *linkList) printList() {
	toPrint := l.head
	for i := 0; i < l.length; i++ {
		fmt.Println(toPrint.value)

		toPrint = toPrint.next

	}
}

// main function
func main() {
	linkList := &linkList{}

	linkList.insertOnFront(100)
	linkList.insertOnFront(200)
	linkList.insertOnTail(300)
	linkList.insertOnFront(400)

	linkList.printList()

	exist := linkList.searchNode(300)
	fmt.Println(exist)

	linkList.deleteNode(300)
	linkList.printList()
}
