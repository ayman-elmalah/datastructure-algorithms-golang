package main

import (
	"fmt"
	"strconv"
)

type Node struct {
	value int
	left  *Node
	right *Node
}

type Tree struct {
	root *Node
}

func (t *Tree) Insert(value int) *Tree {
	node := &Node{
		value: value,
		left:  nil,
		right: nil,
	}

	if t.root == nil {
		t.root = node
		return t
	}

	currentNode := t.root

	for true {
		if value > currentNode.value {
			if currentNode.right == nil {
				currentNode.right = node
				break
			} else {
				currentNode = currentNode.right
			}
		} else {
			if currentNode.left == nil {
				currentNode.left = node
				break
			} else {
				currentNode = currentNode.left
			}
		}
	}

	return t
}

func (t *Tree) Lookup(value int) (string, bool) {
	var path string
	exists := false

	currentNode := t.root
	i := 0

	for currentNode != nil {
		if currentNode.value == value {
			path = path + " > " + strconv.Itoa(i)
			exists = true
			break
		} else if value > currentNode.value {
			currentNode = currentNode.right
		} else {
			currentNode = currentNode.left
		}

		path = path + " > " + strconv.Itoa(i)

		i++
	}

	return path, exists
}

func main() {
	tree := Tree{}
	tree.Insert(50).Insert(150).Insert(200)

	fmt.Println(tree.Lookup(150)) // Output: > 0 > 1 true
}
