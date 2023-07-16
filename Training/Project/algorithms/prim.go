package algorithms

import (
	"container/heap"
	"fmt"
	"math"

	"github.com/falo2/ma/datatypes"
	"github.com/falo2/ma/model"
)

func (g *ExtendedGraph) Prim() (*ExtendedGraph, string) {
	totalWeight := 0.0
	if len(g.Nodes) == 0 {
		return nil, ""
	}

	mst := &ExtendedGraph{Nodes: make(map[int64]*model.Node)}
	used := make(map[int64]bool)
	queue := &datatypes.PriorityQueue{}
	heap.Init(queue)

	heap.Push(queue, model.Edge{Source: 0, Target: 0, Weight: 0})

	for len(*queue) > 0 {
		heapEdge := heap.Pop(queue).(model.Edge)

		if used[heapEdge.Target] {
			continue
		}

		totalWeight += heapEdge.Weight

		used[heapEdge.Target] = true
		mst.Nodes[heapEdge.Target] = &model.Node{ID: heapEdge.Target}

		for _, e := range g.Nodes[heapEdge.Target].Edges {
			if !used[e.Target] {
				heap.Push(queue, *e)
			}
		}

		mst.Nodes[heapEdge.Source].Edges = append(mst.Nodes[heapEdge.Source].Edges, &heapEdge)
		mst.Nodes[heapEdge.Target].Edges = append(mst.Nodes[heapEdge.Target].Edges, &model.Edge{Source: heapEdge.Target, Target: heapEdge.Source, Weight: heapEdge.Weight})
	}

	totalWeight = math.Round(totalWeight*1e6) / 1e6
	return mst, fmt.Sprintf("The MST contains edges with a total weight of %s.", fmt.Sprint(totalWeight))
}
