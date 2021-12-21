// Binary tree
// Traversing

package main

import (
 "fmt"
)

type Node struct {
  Left *Node
  Right *Node
  Data int
}

type BinaryTree struct {
	Root *Node
}

func (n *Node) printInOrder() {
    if n == nil {
    	return
    }
	n.Left.printInOrder()
	fmt.Println(n.Data)
	n.Right.printInOrder()
}       

func main() {
  var myBinarytree BinaryTree
  myBinarytree.Root = &Node {
  	Data: 10,
  }
  myBinarytree.Root.Left = &Node {
  	Data: 5,
  }
  myBinarytree.Root.Left.Left = &Node {
  	Data: 15,
  }
  myBinarytree.Root.Left.Right = &Node {
  	Data: 25,
  }
  myBinarytree.Root.Right = &Node {
  	Data: 7,
  }
  myBinarytree.Root.Right.Left = &Node {
  	Data: 17,
  }
  myBinarytree.Root.Right.Right = &Node {
  	Data: 97,
  }
  myBinarytree.Root.Right.Right.Right = &Node {
  	Data: 100,
  }
  // Traversing
  myBinarytree.Root.printInOrder()

}