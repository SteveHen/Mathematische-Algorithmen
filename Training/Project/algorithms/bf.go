package algorithms

import (
	"fmt"
	"math"

	"github.com/falo2/ma/model"
)

var (
	bestBFWeight = math.Inf(1)
	bestBFPath   []*model.Node
)

func (g *ExtendedGraph) BranchWithoutBound() {
	weight := 0.0
	bestWeight := g.BFHelper(g.Nodes[0], make([]*model.Node, 0), weight)

	fmt.Println("\n= Path =")
	pathString := ""
	for _, node := range bestBFPath {
		pathString += fmt.Sprintf("%d -> ", node.ID)
	}
	fmt.Println(pathString[:len(pathString)-4])

	fmt.Println("\n= Weight =")
	fmt.Println(bestWeight)
}

func (g *ExtendedGraph) BFHelper(node *model.Node, path []*model.Node, weight float64) float64 {
	path = append(path, node)

	if len(path) < len(g.Nodes) {
		for _, edge := range node.Edges {
			if !g.contains(path, edge.Target) {
				g.BFHelper(g.Nodes[edge.Target], path, weight+edge.Weight)
			}
		}
	} else {
		lastEdge, _ := g.getEdge(path[len(path)-1].ID, path[0].ID)
		weight += lastEdge.Weight

		if weight < bestBFWeight {
			bestBFWeight = weight
			bestBFPath = path
		}
	}

	return bestBFWeight
}

func (g *ExtendedGraph) contains(path []*model.Node, target int64) bool {
	for _, node := range path {
		if node.ID == target {
			return true
		}
	}

	return false
}
