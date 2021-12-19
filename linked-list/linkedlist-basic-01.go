// Singly linked list in golang
// 1. struct Node - Data & next
// 2. Set Values
// 3. Prepend
// 4. Append at position
// 5. Delete
// 6. Traversing

package main

import (
	"fmt"
)

type Node struct {
	Data int
	Next *Node
}

type LinkedList struct {
	Head   *Node // Points to the first element in the list
	Length int   // Total elements in list
}

// prepend - append next element in the beginning of Linkedlist
func (l *LinkedList) prepend(node *Node) Node {
	if l.Length > 0 {
		currentHead := l.Head
		node.Next = currentHead
	}
	l.Length += 1

	l.Head = node

	return *node
}

// set - append next element at the end of Linkedlist
func (l *LinkedList) set(node *Node) Node {
	if l.Length == 0 {
		// Add element
		l.Length += 1
		l.Head = node
		return *node
	}
	headToPoint := l.Head
	for headToPoint.Next != nil {
		headToPoint = headToPoint.Next
	}
	headToPoint.Next = node
	l.Length += 1
	return *node
}

// traverse - returns all the the Linkedlist elements
func (l *LinkedList) traverse() []Node {
	var nodes []Node
	if l.Length < 1 {
		return nodes
	}
	elementToPoint := l.Head
	for elementToPoint != nil {
		nodes = append(nodes, *elementToPoint)
		elementToPoint = elementToPoint.Next
	}
	return nodes
}

// delete - delete linkedList element by value
func (l *LinkedList) deleteWithValue(value int) {
	if l.Length < 1 {
		return
	}
	elementToPoint := l.Head
	if elementToPoint.Data == value {
		l.Length -= 1
		l.Head = elementToPoint.Next
		return
	}
	for elementToPoint != nil {
		if elementToPoint.Next == nil {
			break
		}
		if elementToPoint.Next.Data == value {
			nexttoNext := elementToPoint.Next.Next
			elementToPoint.Next = nexttoNext
			l.Length -= 1
		}
		elementToPoint = elementToPoint.Next
	}
	return
}

func main() {
	var myLinkedList LinkedList
	myLinkedList.set(&Node{
		Data: 12,
	}) // 12
	myLinkedList.set(&Node{
		Data: 34,
	}) // 12, 34
	myLinkedList.prepend(&Node{
		Data: 16,
	}) // 16, 12, 34
	myLinkedList.prepend(&Node{
		Data: 45,
	}) // 45, 16, 12, 34
	myLinkedList.prepend(&Node{
		Data: 40,
	}) // 40, 45, 16, 12, 34
	myLinkedList.set(&Node{
		Data: 57,
	}) // 40, 45, 16, 12, 34, 57
	myLinkedList.set(&Node{
		Data: 12,
	}) // 40, 45, 16, 12, 34, 57, 12
	nodes := myLinkedList.traverse()
	fmt.Println("Linked list Nodes", nodes)

	myLinkedList.deleteWithValue(45)  // 40, 16, 12, 34, 57, 12
	myLinkedList.deleteWithValue(100) // 40, 16, 12, 34, 57, 12
	myLinkedList.deleteWithValue(12)  // 40, 16, 34, 57

	nodes = myLinkedList.traverse()
	fmt.Println("Linked list Nodes after Deletion", nodes)
}
