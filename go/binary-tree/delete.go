package binarytree

import "log"

func (n *Node[T]) Delete(data T) {
	if n == nil {
		return
	}

	if data < n.data {
		n.left.Delete(data)
		return
	} else if data > n.data {
		n.right.Delete(data)
		return
	}
	log.Println("on equal ==")
	// on case when node data equal to the given value value
	if n.isLeafNode() {
		log.Println("leaf node")
		n = nil
		return
	}

	if n.left == nil {
		log.Println("no left node", n.data, " << ")
		n.left = n.right
		return
	}

	if n.right == nil {
		log.Println("no right node", n.data, " << ")
		n.right = n.left
		return
	}

	log.Println("when both left and right node exists")
	// on case of when a node has both left and right children
	temp := n.right.FindMax()
	n.data = temp.data
	n.right.Delete(data)

}

func DeleteT[T Any](n *Node[T], data T) *Node[T] {
	if n == nil {
		return &Node[T]{
			left:  nil,
			right: nil,
			data:  data,
		}
	}

	if data < n.data {
		n.left = DeleteT(n.left, data)
	} else if data > n.data {
		n.right = DeleteT(n.right, data)
	} else {
		log.Println("on equal ==")
		// on case when node data equal to the given value value
		// if n.isLeafNode() {
		// 	log.Println("leaf node")
		// 	return DeleteT(n, data)
		// }

		if n.left == nil {
			log.Println("no left node", n.data, " << ")
			return n.right
		}

		if n.right == nil {
			log.Println("no right node", n.data, " << ")
			return n.left
		}

		log.Println("when both left and right node exists")
		// on case of when a node has both left and right children
		temp := n.left.FindMax()
		n.data = temp.data
		n.right = DeleteT(n.right, temp.data)
	}

	return n

}
