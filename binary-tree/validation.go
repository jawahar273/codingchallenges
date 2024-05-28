package binarytree

import (
	"log"
)

func (n *Node[T]) isLeafNode() bool {
	if n.left == nil && n.right == nil {
		return true
	}
	return false
}

func sideValid[T Any](n *Node[T], preRoot *Node[T], topLevelRoot *Node[T]) bool {

	if n == nil {
		return true
	}

	if n.data <= topLevelRoot.data {
		return true
	}

	return false
}

func leftSideValid[T Any](n *Node[T], root *Node[T]) bool {

	if n == nil {
		return true
	}

	if n.data > root.data {
		return false
	}
	// if n.left != nil && n.left.left != nil && n.left.left.data < n.left.data {

	return leftSideValid[T](n.left, root)

}

func rightSideValid[T Any](n *Node[T], root *Node[T]) bool {
	if n == nil {
		return true
	}

	if n.data < root.data {
		return false
	}

	return rightSideValid(n.right, root)
}

func Valid[T Any](n *Node[T]) {
	if n == nil {
		return
	}
	log.Println(
		"validating",
		leftSideValid[T](n.left, n),
		rightSideValid[T](n.right, n),
		// sideValid(n, n, n),
	)
}
