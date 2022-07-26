package main

import "fmt"

const ArraySize int = 10

//Hashtable structure
type HashTable struct {
	array [ArraySize]*bucket
}

//bucket structure
type bucket struct {
	head *bucketNode
}

// bucket node structure
type bucketNode struct {
	key  string
	next *bucketNode
}

//insert into hashmap
func (table *HashTable) Insert(key string) {
	index := hash(key)
	table.array[index].insertIntoBucket(key)
}

//search into hashmap
func (table *HashTable) Search(key string) bool {
	index := hash(key)
	keyExist := table.array[index].searchIntoBucket(key)
	return keyExist
}

//delete into hashmap
func (table *HashTable) Delete(key string) {
	index := hash(key)
	table.array[index].deleteIntoBucket(key)
}

// insert into bucket
func (bucket *bucket) insertIntoBucket(key string) {
	if !bucket.searchIntoBucket(key) {
		newNode := &bucketNode{key: key, next: bucket.head}
		bucket.head = newNode
	} else {
		fmt.Println("key already exists")
	}
}

// search into bucket
func (bucket *bucket) searchIntoBucket(key string) bool {
	for currentNode := bucket.head; currentNode != nil; currentNode = currentNode.next {
		if currentNode.key == key {
			return true
		}
	}
	return false
}

// delete into bucket
func (bucket *bucket) deleteIntoBucket(key string) {
	if bucket.head.key == key {
		bucket.head = bucket.head.next
		return
	}
	previousNode := bucket.head

	for previousNode.next != nil {
		if previousNode.next.key == key {
			previousNode.next = previousNode.next.next
		}
		previousNode = previousNode.next
	}
}

// hash
func hash(key string) int {
	sum := 0
	for _, i := range key {
		sum += int(i)
	}
	return sum % ArraySize
}

// init
func Init() *HashTable {
	table := &HashTable{}

	for i := range table.array {
		table.array[i] = &bucket{}
	}

	return table
}

func main() {
	testHashMap := Init()

	testHashMap.Insert("test")
	testHashMap.Insert("ali")
	testHashMap.Insert("ahmad")

	// testHashMap.Insert(testBucket)
	fmt.Println(*testHashMap)

	res := testHashMap.Search("ali")
	fmt.Println(res)

	testHashMap.Delete("ali")

}
