package main

import (
	"fmt"
)

type Node struct {
	val         int
	left, right *Node
}

type BST struct {
	root *Node
}

func (bst *BST) insert(v int) {
	n := &Node{v, nil, nil}

	if bst.root == nil {
		bst.root = n
		return
	}

	cur := bst.root

	for cur != nil {
		if n.val < cur.val {
			if cur.left == nil {
				cur.left = n
				break
			} else {
				cur = cur.left
			}
		} else {
			if cur.right == nil {
				cur.right = n
				break
			} else {
				cur = cur.right
			}
		}
	}
}

func (bst *BST) preOrder(root *Node) {
	if root != nil {
		fmt.Printf("%v ", root.val)
		bst.preOrder(root.left)
		bst.preOrder(root.right)
	}
}

func (bst *BST) inOrder(root *Node) {
	if root != nil {
		bst.preOrder(root.left)
		fmt.Printf("%v ", root.val)
		bst.preOrder(root.right)
	}
}

func (bst *BST) postOrder(root *Node) {
	if root != nil {
		bst.preOrder(root.left)
		bst.preOrder(root.right)
		fmt.Printf("%v ", root.val)
	}
}

func main() {
	bst := new(BST)

	bst.insert(8)
	bst.insert(4)
	bst.insert(3)
	bst.insert(1)
	bst.insert(10)
	bst.insert(40)
	bst.insert(21)

	fmt.Println("root node is ", bst.root)

	bst.preOrder(bst.root)
	fmt.Println()
	bst.inOrder(bst.root)
	fmt.Println()
	bst.postOrder(bst.root)
}
