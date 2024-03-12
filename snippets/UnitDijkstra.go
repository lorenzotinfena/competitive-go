package _

import (
	"github.com/lorenzotinfena/goji/collections/graph"
)

func UnitDijkstra() {
	var _n, _start, _end int
	var _adj [][]int

	dijkstranodes := make([]bool, _n)
	it := graph.UnitDijkstra(
		_start,
		func(i int) []int { return _adj[i] },
		func(i int) { dijkstranodes[i] = true },
		func(i int) bool {
			return dijkstranodes[i]
		},
	)
	path := []int{-1}
	for it.HasNext() {
		tmp := it.Next()
		if tmp.Vertex == _end {
			path = []int{}
			for tmp != nil {
				path = append(path, tmp.Vertex+1)
				tmp = tmp.Previous
			}
			for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
				path[i], path[j] = path[j], path[i]
			}
			break
		}
	}
}
