package algorithms

import (
	"math"

	"github.com/falo2/ma/model"
)

func (g *ExtendedGraph) MooreBellmanFord(source, target *model.Node, verbose bool) ([]model.Edge, map[int64]*model.Node, float64) {
	dist := make(map[int64]float64)
	previous := make(map[int64]*model.Node)
	previous[source.ID] = source

	minCapacity := math.Inf(1)

	for i := 0; i < len(g.Nodes); i++ {
		dist[int64(i)] = math.Inf(1)
	}

	dist[source.ID] = 0

	edgeCount := 0
	for i := 1; i < len(g.Nodes); i++ {
		for j := 0; j < len(g.Nodes); j++ {
			for k := 0; k < len(g.Nodes[int64(j)].Edges); k++ {
				edge := g.Nodes[int64(j)].Edges[k]
				if edge.Capacity > 0 {
					nodeCurrent := g.Nodes[edge.Source]
					nodeNext := g.Nodes[edge.Target]

					if dist[nodeCurrent.ID]+edge.Cost < dist[nodeNext.ID] {
						edgeCount++
						dist[nodeNext.ID] = dist[nodeCurrent.ID] + edge.Cost
						previous[nodeNext.ID] = nodeCurrent
					}
				}
			}
		}
	}

	for i := 0; i < len(g.Nodes); i++ {
		for j := 0; j < len(g.Nodes[int64(i)].Edges); j++ {
			edge := g.Nodes[int64(i)].Edges[j]
			if edge.Capacity > 0 {
				if dist[edge.Source]+edge.Cost < dist[edge.Target] {
					currentNode := g.Nodes[edge.Source]
					for j := 0; j < len(g.Nodes); j++ {
						currentNode = previous[currentNode.ID]
					}

					cycle := make([]model.Edge, 0)
					start := currentNode.ID
					for previous[currentNode.ID].ID != start {
            cycle, minCapacity = g.addCycleEdge(previous[currentNode.ID].ID, currentNode.ID, cycle, minCapacity)
						currentNode = previous[currentNode.ID]
					}

					cycle, minCapacity = g.addCycleEdge(previous[currentNode.ID].ID, currentNode.ID, cycle, minCapacity)
					return cycle, previous, minCapacity
				}
			}
		}
	}

	if target != nil {
		minCapacity = math.Round(dist[target.ID]*1e6) / 1e6
	}
	return nil, previous, minCapacity
}

func (g *ExtendedGraph) addCycleEdge(source, target int64, cycle []model.Edge, minCapacity float64) ([]model.Edge, float64) {
	cycleEdge, _ := g.getEdge(source, target)
	cycle = append(cycle, *cycleEdge)

	minCapacity = math.Min(minCapacity, cycleEdge.Capacity)
	return cycle, minCapacity
}
