// Circular singly linked list
// 1. Add element at the beginning
// 2. Add element at the end
// 3. Add element at the given position
// 4. Traversing
// 5. Deleting

package main

import (
  "fmt"
)

type Node struct {
	Data int
	Next *Node
}

type CircularSinglyLinkedList struct {
	Last *Node
	Length int
}

// set - add node at the end of the list
func (c *CircularSinglyLinkedList) set(node *Node) Node {
  lastItem := c.Last
	if c.Length == 0 {
		node.Next = node
	} else {
    node.Next = lastItem.Next // First element
    lastItem.Next = node
  }
	c.Length += 1
	c.Last = node
  return *node
}

// set - add node at the end of the list
// Position starts with 0
func (c *CircularSinglyLinkedList) setAtPosition(node *Node, position int) Node {
  if position > c.Length {
    return Node{}
  }
  if position == 0 {
    return c.prepend(node)
  }
  if position == c.Length {
    return c.set(node)
  }
  elementPosition := 0
  startElement := c.Last.Next
  for elementPosition < c.Length {
    if (position - 1) == elementPosition {
      node.Next = startElement.Next
      startElement.Next = node
      c.Length += 1
      break
    }
    startElement = startElement.Next
    elementPosition += 1
  }

  return *node
}

// prepend - add node at the beginning of the list
func (c *CircularSinglyLinkedList) prepend(node *Node) Node {
  lastItem := c.Last
  if c.Length == 0 {
    node.Next = node
    c.Last = node
  } else {
    node.Next = lastItem.Next // First element
    lastItem.Next = node
  }
  c.Length += 1
  return *node
}

// traversing - traverse through all the list nodes
func (c *CircularSinglyLinkedList) traversing() []Node {
  var nodes []Node
  if c.Length < 1 {
    return nodes
  }
  startingElement := c.Last.Next
  itemsCount := c.Length
  for itemsCount > 0 {
    nodes = append(nodes, *startingElement)
    startingElement = startingElement.Next
    itemsCount -=1
  }
  return nodes
}

// deleteWithValue - delete all the nodes with given value
func (c *CircularSinglyLinkedList) deleteWithValue(value int) {
  if c.Length < 1 {
    return
  }
  startingElement := c.Last
  itemsCount := c.Length

  for itemsCount > 0 {
    if startingElement.Next.Data == value {
      if c.Length <= 1 {
        c.Length -= 1
        c.Last = nil
        break
      }
      if c.Last == startingElement.Next {
        c.Last = startingElement
      }
      startingElement.Next = startingElement.Next.Next
      c.Length -= 1
    } else {
      startingElement = startingElement.Next // Increment if next element if the value not found
    }
    itemsCount -=1
  }
  return
}

func main() {
  var myCircularSinglyLinkedList CircularSinglyLinkedList
  myCircularSinglyLinkedList.set(&Node {
  		Data: 12,
  })                                                // 12
  myCircularSinglyLinkedList.set(&Node{
  		Data: 24,
  })                                                // 12, 24
  myCircularSinglyLinkedList.set(&Node{
      Data: 30,
  })                                                // 12, 24, 30
  myCircularSinglyLinkedList.prepend(&Node{
      Data: 77,
  })                                                // 77, 12, 24, 30
  myCircularSinglyLinkedList.set(&Node{
      Data: 57,
  })                                                // 77, 12, 24, 30, 57
  myCircularSinglyLinkedList.set(&Node{
      Data: 57,
  })                                                // 77, 12, 24, 30, 57, 57
  myCircularSinglyLinkedList.set(&Node{
      Data: 57,
  })                                                // 77, 12, 24, 30, 57, 57, 57
  myCircularSinglyLinkedList.setAtPosition(&Node{
      Data: 1,
  }, 0)                                             // 1, 77, 12, 24, 30, 57, 57, 57
  myCircularSinglyLinkedList.setAtPosition(&Node{
      Data: 3,
  }, 3)                                             // 1, 77, 12, 3, 24, 30, 57, 57, 57
  myCircularSinglyLinkedList.setAtPosition(&Node{
      Data: 9,
  }, 9)                                             // 1, 77, 12, 3, 24, 30, 57, 57, 57, 9
  myCircularSinglyLinkedList.setAtPosition(&Node{
      Data: 10,
  }, 11)                                            // 1, 77, 12, 3, 24, 30, 57, 57, 57, 9
  nodes := myCircularSinglyLinkedList.traversing()
  fmt.Println("\nLinked list Nodes \n", nodes)

  myCircularSinglyLinkedList.deleteWithValue(88)
  myCircularSinglyLinkedList.deleteWithValue(57)

  nodes = myCircularSinglyLinkedList.traversing()
  fmt.Println("\nLinked list Nodes after deletion\n", nodes)
}