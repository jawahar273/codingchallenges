package binarytree

import (
	"log"
	"testing"
)

func TestDeleteTree(t *testing.T) {
	tree := NewNode(50)
	tree.BulkInsert([]int{20, 100, 10, 30, 1})

	deleteValue := 20

	InOrderDisplay(tree)

	// tree.Delete(deleteValue)
	tree = DeleteT(tree, deleteValue)

	t.Log(
		tree.Find(deleteValue),
	)
	InOrderDisplay(tree)

}

func TestTree(t *testing.T) {

	// tree := NewNode(50)
	// tree.BulkInsert([]int{20, 100, 10, 30, 1})

	tree := NewNode(15)
	tree.BulkInsert([]int{54, 13, 62, 35})

	log.Println(" ========= pre order =========")
	PreOrderDisplay(tree)
	log.Println(" ========= in order =========")
	InOrderDisplay(tree)

	// tree = NewNode(5)
	// tree.BulkInsert([]int{3, 8, 9, 10})
	// t.Log(
	// 	CheckBalance[int](tree),
	// )

	// tree = NewNode(5)
	// tree.BulkInsert([]int{3, 8, 1, 4})
	// t.Log(
	// 	CheckBalance[int](tree),
	// )

	tree = NewNode(5)
	tree.BulkInsert([]int{8, 5, 6, 4, 7, 1001})
	Valid[int](tree)
	t.Log("==>>>>>",
		tree.FindMax().data,
		tree.FindMin().data,
	)

	// tree = NewNode(5)
	// tree.BulkInsert([]int{5, 8, 20})
	// Valid[int](tree)

	// tree = NewNode(5)
	// left := NewNode(5)
	// right := NewNode(8)
	// invalid := NewNode(20)
	// tree.left = left
	// tree.right = right
	// tree.left.right = invalid
	// Valid(tree)

}
