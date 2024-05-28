package binarytree

type Any interface {
	int | float64
}

type Node[T Any] struct {
	data  T
	left  *Node[T]
	right *Node[T]
}

func NewNode[T Any](data T) *Node[T] {
	return &Node[T]{data: data, left: nil, right: nil}
}

func (n *Node[T]) Insert(data T) {

	if n == nil {
		return
	}

	if n.data > data {
		// log.Println("left side check", n.data)
		if n.left == nil {
			// log.Println("to left node", data)
			n.left = NewNode(data)
		}
		n.left.Insert(data)
		return
	} else if n.data < data {
		// log.Println("right side check", n.data)
		if n.right == nil {
			// log.Println("to right node", data)
			n.right = NewNode(data)
		}
		n.right.Insert(data)
		return
	} else if n.data == data {
		// log.Println("right(equal) node", data)
		n.right.Insert(data)
		return
	}
}

func (n *Node[T]) BulkInsert(items []T) {
	if n == nil {
		return
	}

	for _, item := range items {
		n.Insert(item)
	}
}

func (n *Node[T]) Find(value T) bool {
	if n == nil {
		return false
	}

	if n.data == value {
		return true
	}

	if value < n.data {
		return n.left.Find(value)
	}

	return n.right.Find(value)

}

func (n *Node[T]) FindMax() *Node[T] {
	if n == nil {
		return nil
	}

	if n.right != nil {
		return n.right.FindMax()
	}

	return n
}

func (n *Node[T]) FindMin() *Node[T] {
	if n == nil {
		return nil
	}

	if n.left != nil {
		return n.left.FindMin()
	}

	return n
}
