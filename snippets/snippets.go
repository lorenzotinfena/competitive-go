package _

import (
	"github.com/lorenzotinfena/goji/collections"
	"github.com/lorenzotinfena/goji/collections/graph"
	"github.com/lorenzotinfena/goji/collections/heap"
)

func main() {
	// #region WeightedDijkstra
	var n, start, end int
	var adj [][]collections.Pair[int, int]

	dijkstranodes := make([]*graph.DijkstraNode[int, int], n)
	fibonaccihelper := make([][]*heap.FibonacciHeapNode[int], n)
	it := graph.WeightedDijkstra(
		start,
		func(i int) []collections.Pair[int, int] { return adj[i] },
		func(i int, dn *graph.DijkstraNode[int, int]) { dijkstranodes[i] = dn },
		func(i int) *graph.DijkstraNode[int, int] {
			return dijkstranodes[i]
		},
		func(i int) bool { return dijkstranodes[i] != nil },
		func(i int, fhn []*heap.FibonacciHeapNode[int]) { fibonaccihelper[i] = fhn },
		func(i int) []*heap.FibonacciHeapNode[int] { return fibonaccihelper[i] },
		func(i int) bool { return fibonaccihelper[i] != nil },
		func(i int) { fibonaccihelper[i] = nil },
	)
	path := []int{-1}
	for it.HasNext() {
		tmp := it.Next()
		if tmp.Vertex == end {
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
	// #endregion
}
