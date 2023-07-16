package algorithms

import (
	//"fmt"
	"math"

	"github.com/falo2/ma/model"
)

func (g *ExtendedGraph) EdmondsKarp(source, target int64) (*ExtendedGraph, float64) {
	flowGraph := g.CopyFullGraph()
	flowGraph.checkResidualFlowGraph()

  path, _ := flowGraph.BFS(source, target, nil)

	var edge *model.Edge
	var flowGraphFlow float64

	for len(path) > 0 {
		maxFlow := math.MaxFloat64

		for i := 0; i < len(path)-1; i++ {
			current := path[i]
			next := path[i+1]

			edge, _ = flowGraph.getEdge(current.ID, next.ID)
			if edge.Capacity < maxFlow {
				maxFlow = edge.Capacity
			}
		}

		for i := 0; i < len(path)-1; i++ {
			current := path[i]
			next := path[i+1]

			edge, _ = flowGraph.getEdge(current.ID, next.ID)
			edge.Weight += maxFlow
			edge.Capacity -= maxFlow

			edge, _ = flowGraph.getEdge(next.ID, current.ID)
			edge.Weight -= maxFlow
			edge.Capacity += maxFlow
		}

		flowGraphFlow += maxFlow
		path, _ = flowGraph.BFS(source, target, nil)
	}

	return flowGraph, flowGraphFlow
}

func (g *ExtendedGraph) checkResidualFlowGraph() {
	for i := 0; i < len(g.Nodes); i++ {
		for j := 0; j < len(g.Nodes[int64(i)].Edges); j++ {
			edge := g.Nodes[int64(i)].Edges[j]
			residualEdge, _ := g.getEdge(edge.Target, edge.Source)
			if residualEdge == nil {
				residualEdge = &model.Edge{
					Source:   edge.Target,
					Target:   edge.Source,
					Weight:   0,
					Capacity: 0,
					Cost:     -edge.Cost,
					Residual: true,
				}

				g.Nodes[edge.Target].Edges = append(g.Nodes[edge.Target].Edges, residualEdge)
			}
		}
	}
}
