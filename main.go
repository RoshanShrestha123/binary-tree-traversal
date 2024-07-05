package main

import "fmt"

type Node struct {
	value int
	left  *Node
	right *Node
}

func (n *Node) traverse() {

	if n == nil {
		return
	}

	val := n.value
	fmt.Printf("Value: %d\n", val)
	n.left.traverse()
	n.right.traverse()
}

func main() {
	fmt.Println("Binary Tree Traversal")
	two := &Node{value: 2}
	four := &Node{value: 4}
	five := &Node{value: 5}
	three := &Node{value: 3, left: four, right: five}
	one := &Node{value: 1, left: two, right: three}

	one.traverse()

}
