// Doubly linked list
// Add element at the start
// Add element at the end
// Add at given position
// Delete based on the given value

package main

import (
  "fmt"
)

type Node struct {
	Data int
	Next *Node
	Prev *Node
}

type DoublyLinkedList struct {
	Head * Node
	Length int
}

func (d *DoublyLinkedList) set(node *Node) Node {
	if d.Length < 1 {
		d.Head = node
		d.Length += 1
		return *node
	}
	itemToPoint := d.Head
	if itemToPoint.Next != nil {
		for itemToPoint != nil {
			if itemToPoint.Next == nil {
				break
			}
			itemToPoint = itemToPoint.Next
		}
	}
	node.Prev = itemToPoint
	itemToPoint.Next = node
	d.Length += 1
	return *node
}

func (d *DoublyLinkedList) setAtPosition(node *Node, pos int) Node {
	if pos > d.Length {
      return Node{}
	}
	if pos == 0 {
		return d.prepend(node)
	}
	if pos == d.Length {
		return d.set(node)
	}
	currentIndex := 0
	currentItem := d.Head
	for currentItem != nil {
		if currentIndex == (pos-1) {
		   nextNode := currentItem.Next
           node.Next = nextNode
           node.Prev = currentItem
           currentItem.Next = node
           nextNode.Prev = node
           d.Length += 1
           break
		}
		currentItem = currentItem.Next
		currentIndex += 1
	}
	return *node
}

func (d *DoublyLinkedList) prepend(node *Node) Node {
	if d.Length < 1 {
		d.Head = node
		d.Length += 1
		return *node
	}
	itemToPoint := d.Head

	itemToPoint.Prev = node
	node.Next = itemToPoint
	d.Head = node
	d.Length += 1
	return *node
}

func (d *DoublyLinkedList) traverse() []Node {
    var nodes []Node
    itemToPoint := d.Head
    for itemToPoint != nil {
    	nodes = append(nodes, *itemToPoint)
    	itemToPoint = itemToPoint.Next
    }

	return nodes
}

func (d *DoublyLinkedList) delete(value int) { 
    if d.Length < 1 {
    	return
    }
    pointToElement := d.Head

    for pointToElement != nil {
    	if pointToElement.Data == value {
            if pointToElement.Prev != nil {
           	  pointToElement.Prev.Next = pointToElement.Next
            } else {
            	d.Head = pointToElement.Next
            }
           
            if pointToElement.Next != nil {
           	  pointToElement.Next.Prev = pointToElement.Prev
            }
           
           d.Length -= 1
    	}
    	pointToElement = pointToElement.Next
    }
	return
}

func main() {
  var myDoublyLinkedList DoublyLinkedList
  myDoublyLinkedList.set(&Node{
  	Data: 10,
  })                                             // 10
  myDoublyLinkedList.set(&Node{
  	Data: 20,
  })                                             // 10, 20
  myDoublyLinkedList.prepend(&Node{
  	Data: 5,
  })                                             // 5, 10, 20
  myDoublyLinkedList.prepend(&Node{
  	Data: 1,
  })                                             // 1, 5, 10, 20
  myDoublyLinkedList.setAtPosition(&Node{
  	Data: 3,
  }, 2)                                          // 1, 5, 3, 10, 20
  myDoublyLinkedList.setAtPosition(&Node{
  	Data: 50,
  }, 5)                                          // 1, 5, 3, 10, 20, 50
  myDoublyLinkedList.setAtPosition(&Node{
  	Data: 50,
  }, 6)                                          // 1, 5, 3, 10, 20, 50, 50
  myDoublyLinkedList.setAtPosition(&Node{
  	Data: 60,
  }, 60)                                         // 1, 5, 3, 10, 20, 50, 50                                    
  nodes := myDoublyLinkedList.traverse()
  fmt.Println("Doubly Linked List nodes", nodes)

  myDoublyLinkedList.delete(1)
  myDoublyLinkedList.delete(50)

  nodes = myDoublyLinkedList.traverse()
  fmt.Println("Doubly Linked List nodes after deletion", nodes)
}