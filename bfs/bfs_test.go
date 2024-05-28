package bfs

import (
	"fmt"
	"testing"
)

func TestBfs(t *testing.T) {
	one := &Node[int]{id: 1, visited: false, adjacentGraph: make(map[int]*Node[int])}
	two := &Node[int]{id: 2, visited: false, adjacentGraph: make(map[int]*Node[int])}
	three := &Node[int]{id: 3, visited: false, adjacentGraph: make(map[int]*Node[int])}
	four := &Node[int]{id: 4, visited: false, adjacentGraph: make(map[int]*Node[int])}
	five := &Node[int]{id: 5, visited: false, adjacentGraph: make(map[int]*Node[int])}

	// g := Graph[int]{}

	// g.Holder[0] = []*Node[int]{}
	// g.Holder[1] = []*Node[int]{}
	// g.Holder[3] = []*Node[int]{}
	// g.Holder[4] = []*Node[int]{}
	// g.Holder[5] = []*Node[int]{}

	two.adjacentGraph[one.id] = one
	three.adjacentGraph[two.id] = two
	one.adjacentGraph[three.id] = three
	three.adjacentGraph[four.id] = four
	four.adjacentGraph[five.id] = five

	root := &Node[int]{id: 0, adjacentGraph: map[int]*Node[int]{}}
	root.adjacentGraph[one.id] = one
	root.adjacentGraph[two.id] = two
	root.adjacentGraph[three.id] = three

	temp := Bfs(two)
	for _, n := range temp {
		fmt.Println("-->", n.id)
	}
	// Display(root)
}
