package main

import "fmt"

type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

func (n *Node) insert(k int) {
	if n.Key < k {
		// move right
		if n.Right == nil {
			n.Right = &Node{Key: k}
		} else {
			n.Right.insert(k)
		}
	} else if n.Key > k {
		// move left
		if n.Left == nil {
			n.Left = &Node{Key: k}
		} else {
			n.Left.insert(k)
		}
	}
}

func (n *Node) search(k int) bool {
	if n == nil {
		return false
	}
	if n.Key < k {
		// move right
		return n.Right.search(k)
	} else if n.Key > k {
		// move left
		return n.Left.search(k)
	}
	return true
}

func main() {
	tree := &Node{Key: 100}
	tree.insert(50)
	tree.insert(69)
	tree.insert(420)
	tree.insert(33)
	fmt.Println(tree.search(69))
}
