package bfs

import (
	"fmt"

	"cc.fyi/queue"
	"cc.fyi/types"
)

type Node[T types.AnyNumber] struct {
	id            T
	visited       bool
	adjacentGraph map[T]*Node[T]
}

type Graph[T types.AnyNumber] struct {
	Holder map[T][]*Node[T]
}

func Display[T types.AnyNumber](root *Node[T]) {
	if root == nil {
		return
	}
	temp := root
	for _, network := range temp.adjacentGraph {
		if !network.visited {
			fmt.Println(network.id)
		}
	}
}

// ref:- https://cybernetist.com/2019/03/09/breadth-first-search-using-go-standard-library/
// ref:- https://www.cs.usfca.edu/~galles/visualization/BFS.html
func Bfs[T types.AnyNumber](root *Node[T]) []*Node[T] {
	q := queue.Queue[*Node[T]]{}
	root.visited = true
	q.Enqueue(root)
	var result []*Node[T]
	for !q.IsEmpty() {
		v := q.Dequeue()
		for _, network := range v.adjacentGraph {
			_, ok := v.adjacentGraph[network.id]
			if ok && !network.visited {
				network.visited = true
				q.Enqueue(network)
				// fmt.Println("bfs:", network.id)
				result = append(result, network)
			}
		}
	}
	return result
}
