package binarytree

import "math"

func checkBalanceRec[T Any](n *Node[T]) int64 {
	if n == nil {
		return 0
	}

	leftHeight := checkBalanceRec[T](n.left)
	if leftHeight == -1 {
		return -1
	}

	rightHeight := checkBalanceRec[T](n.right)
	if rightHeight == -1 {
		return -1
	}

	diff := math.Abs(float64(leftHeight - rightHeight))
	if diff > 1 {
		return -1
	}

	return 1 + max(leftHeight, rightHeight)
}

func CheckBalance[T Any](n *Node[T]) bool {
	return checkBalanceRec[T](n) != -1
}
