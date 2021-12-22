// Binary search tree
// Insert new item
// traversing

package main

import (
	"fmt"
)

type Node struct {
	Data int
    Left *Node
    Right *Node
}

type BinarySearchTree struct {
  Root *Node
}

func (b *BinarySearchTree) insert(node *Node) Node {
	if b.Root == nil {
		b.Root = node
		return *node
	}
	currentNode := b.Root
    currentNode.getNode(node)
    return *node
}

func (node *Node) getNode(newNode *Node) {
  if newNode.Data < node.Data {
  	if node.Left == nil {
	  	node.Left = newNode
	  	fmt.Printf("Addint newNode %v to left of node %v\n",newNode.Data, node.Data )
	  	return
	  } else {
	  	node.Left.getNode(newNode)
	  }
  } else {
  	if node.Right == nil {
	  	node.Right = newNode
	  	fmt.Printf("Addint newNode %v to Right of node %v\n", newNode.Data, node.Data )
	  	return
	} else {
		node.Right.getNode(newNode)
	}
  }
}

func (n *Node) traverse() {
	if n == nil {
		return
	}
	n.Left.traverse()
	fmt.Println(n.Data)
	n.Right.traverse()
}

func main() {
  var myBinarySearchTree BinarySearchTree
  myBinarySearchTree.insert(&Node{Data: 50})
  myBinarySearchTree.insert(&Node{Data: 90})
  myBinarySearchTree.insert(&Node{Data: 30})
  myBinarySearchTree.insert(&Node{Data: 20})
  myBinarySearchTree.insert(&Node{Data: 60})
  myBinarySearchTree.insert(&Node{Data: 100})
  myBinarySearchTree.insert(&Node{Data: 10})
  fmt.Println("Traversing")
  myBinarySearchTree.Root.traverse()
}