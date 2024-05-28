package binarytree

import "log"

func InOrderDisplay[T Any](root *Node[T]) {
	if root == nil {
		// log.Fatal("nil tree")
		return
	}
	InOrderDisplay(root.left)
	log.Println(root.data)
	InOrderDisplay(root.right)
}

func PreOrderDisplay[T Any](root *Node[T]) {
	if root == nil {
		// log.Fatal("nil tree")
		return
	}
	log.Println(root.data)
	PreOrderDisplay(root.left)
	PreOrderDisplay(root.right)
}

func PostOrder[T Any](root *Node[T]) {
	if root == nil {
		// log.Fatal("nil tree")
		return
	}
	PostOrder(root.left)
	PostOrder(root.right)
	log.Println(root.data)
}
