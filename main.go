package main

import "fmt"

type node struct {
	data int
	next *node
}

type linkedList struct {
	length int
	head   *node
	tail   *node
}

// Returns the length of linked list
func (list linkedList) getLength() int {
	return list.length
}

// Prints the elements of linked list
func (list linkedList) printList() {
	for list.head != nil {
		fmt.Printf("%v -> ", list.head.data)
		list.head = list.head.next
	}

	fmt.Println()
}

// Adds new node in the linked list
func (list *linkedList) pushBack(node *node) {
	if list.head == nil {
		list.head = node
		list.tail = node
		list.length++

		return
	}

	list.tail.next = node
	list.tail = node
	list.length++
}

// Removes duplicates elements.
// This is the best solution in runtime complexity which is O(n)
// also in space complexity is O(n) because we are creating a hash set.
func (list *linkedList) deleteDuplicates() {
	hashSet := map[int]bool{}
	var previous *node
	temp := list.head

	for temp != nil {
		if hashSet[temp.data] {
			previous.next = temp.next
		} else {
			hashSet[temp.data] = true
			previous = temp
		}

		temp = temp.next
	}
}

// Removes duplicates without using a storage.
// This is not the best solution in runtime complexity which is O(n^2); however,
// we reduce the space complexity to O(1) because we are not using a extra data structure
func (list *linkedList) deleteDuplicatesWithoutBuffer() {
	current := list.head

	for current != nil {
		runner := current

		for runner.next != nil {
			if runner.next.data == current.data {
				runner.next = runner.next.next
			} else {
				runner = runner.next
			}
		}

		current = current.next
	}
}

func main() {
	node1 := &node{data: 10}
	node2 := &node{data: 3}
	node3 := &node{data: 11}
	node4 := &node{data: 11}
	linkedList1 := linkedList{}

	linkedList1.pushBack(node1)
	linkedList1.pushBack(node2)
	linkedList1.pushBack(node3)
	linkedList1.pushBack(node4)

	linkedList1.deleteDuplicatesWithoutBuffer()
	linkedList1.printList()
}
