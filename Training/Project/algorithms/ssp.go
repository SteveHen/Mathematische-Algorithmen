package algorithms

import (
	"math"

	"github.com/falo2/ma/model"
)

func (g *ExtendedGraph) SuccessiveShortestPath(balance map[int64]float64, verbose bool) (*ExtendedGraph, float64) {
	result := g.CopyFullGraph()
	result.checkResidualFlowGraph()
	currentBalance := make(map[int64]float64)

	for i := 0; i < len(result.Nodes); i++ {
		for j := 0; j < len(result.Nodes[int64(i)].Edges); j++ {
			edge := result.Nodes[int64(i)].Edges[j]
			if !edge.Residual && edge.Cost < 0 {
				edge.Weight = edge.Capacity
				edge.Capacity = 0

				revEdge, _ := result.getEdge(edge.Target, edge.Source)
				revEdge.Capacity = edge.Weight
			}

			currentNode := edge.Source
			currentBalance[currentNode] += edge.Weight

			nextNode := edge.Target
			currentBalance[nextNode] -= edge.Weight
		}
	}

	iteration := 0

	for {
		var source *model.Node
		for i := 0; i < len(result.Nodes); i++ {
			if balance[int64(i)]-currentBalance[int64(i)] > 0 {
				source = result.Nodes[int64(i)]
				break
			}
		}

		var target *model.Node
		if source == nil {
			target = nil
		} else {
			visited := make(map[int64]bool)
			_, visited = result.BFS(source.ID, -1, visited)
			for i := 0; i < len(result.Nodes); i++ {
				current := result.Nodes[int64(i)]
				if visited[current.ID] && balance[current.ID]-currentBalance[current.ID] < 0 {
					target = current
					break
				}
			}
		}

		if source == nil || target == nil {
			if result.isCostMinimal(balance, currentBalance) {
				return result, result.calculateMinCosts()
			} else {
				return nil, math.Inf(1)
			}
		}

		_, prev, _ := result.MooreBellmanFord(source, nil, verbose)

		var path []*model.Node
		current := target
		for current.ID != source.ID {
			path = append([]*model.Node{current}, path...)
			current = prev[current.ID]
		}
		path = append([]*model.Node{source}, path...)

		gamma := math.Min(balance[source.ID]-currentBalance[source.ID], currentBalance[target.ID]-balance[target.ID])

		for i := 0; i < len(path)-1; i++ {
			currentNode := path[i]
			nextNode := path[i+1]

			edge, _ := result.getEdge(currentNode.ID, nextNode.ID)
			gamma = math.Min(gamma, edge.Capacity)
		}

		for i := 0; i < len(path)-1; i++ {
			currentNode := path[i]
			nextNode := path[i+1]

			edge, _ := result.getEdge(currentNode.ID, nextNode.ID)
			revEdge, _ := result.getEdge(nextNode.ID, currentNode.ID)

			edge.Weight += gamma
			edge.Capacity -= gamma

			revEdge.Weight -= gamma
			revEdge.Capacity += gamma
		}

		currentBalance[source.ID] += gamma
		currentBalance[target.ID] -= gamma
		iteration++
	}
}

func (g *ExtendedGraph) isCostMinimal(balance map[int64]float64, currentBalance map[int64]float64) bool {
	for i := 0; i < len(g.Nodes); i++ {
		if balance[int64(i)] != currentBalance[int64(i)] {
			return false
		}
	}

	return true
}
