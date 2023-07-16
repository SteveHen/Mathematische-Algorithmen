package algorithms

import (
	"fmt"
	//"strconv"

	"github.com/falo2/ma/model"
)

func (g *ExtendedGraph) CycleCanceling(balance map[int64]float64, verbose bool) (*ExtendedGraph, float64) {
	result := g.CopyFullGraph()

	source := int64(len(result.Nodes))
	target := int64(len(result.Nodes) + 1)
	result.Nodes[source] = &model.Node{ID: source}
	result.Nodes[target] = &model.Node{ID: target}

	for node := range result.Nodes {
		if node != source && node != target {
			balanceNode := balance[node]
			if balanceNode > 0 {
				result.Nodes[source].Edges = append(result.Nodes[source].Edges, &model.Edge{
					Source:   source,
					Target:   node,
					Capacity: balanceNode,
				})
			} else if balanceNode < 0 {
				result.Nodes[node].Edges = append(result.Nodes[node].Edges, &model.Edge{
					Source:   node,
					Target:   target,
					Capacity: -balanceNode,
				})
			}
		}
	}

	flowGraph, maxFlow := result.EdmondsKarp(source, target)

	balanceSum := 0.0
	for _, value := range balance {
		if value > 0 {
			balanceSum += value
		}
	}

	if verbose {
		fmt.Println("Max flow: ", maxFlow)
		fmt.Println("Balance sum: ", balanceSum)
	}

	if maxFlow != balanceSum {
		return nil, 0
	}

	var ready bool
  visited := make(map[int64]bool)

	for !ready {
		for i := 0; i < len(flowGraph.Nodes); i++ {
			visited[int64(i)] = true

			cycle, _, cycleMinCapacity := flowGraph.MooreBellmanFord(flowGraph.Nodes[int64(i)], nil, verbose)

			var visitedNodes int
			for _, visit := range visited {
				if visit {
					visitedNodes++
				}
			}

			if visitedNodes == len(flowGraph.Nodes) {
				ready = true
				break
			}

			for _, edge := range cycle {
				flowEdge, _ := flowGraph.getEdge(edge.Source, edge.Target)
				flowEdge.Capacity -= cycleMinCapacity
				flowEdge.Weight += cycleMinCapacity

				revEdge, _ := flowGraph.getEdge(edge.Target, edge.Source)
				revEdge.Capacity += cycleMinCapacity
				revEdge.Weight -= cycleMinCapacity
			}
		}
	}

	cost := flowGraph.calculateMinCosts()
	return flowGraph, cost
}

func (g *ExtendedGraph) calculateMinCosts() float64{
	cost := 0.0
	for _, node := range g.Nodes {
		for _, edge := range node.Edges {
			if !edge.Residual {
				cost += edge.Weight * edge.Cost
			}
		}
	}

	return cost
}
