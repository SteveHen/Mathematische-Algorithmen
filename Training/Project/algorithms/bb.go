package algorithms

import (
	"fmt"
	"math"

	"github.com/falo2/ma/model"
)

var bestBBWeight *float64

func (g *ExtendedGraph) BranchBound() {
	initBestWeight := math.Inf(1)
	bestBBWeight = &initBestWeight
	g.BBHelper(g.Nodes[0], make([]*model.Node, 0), 0.0)
	fmt.Println("\n= Weight =")
	fmt.Println(*bestBBWeight)
}

func (g *ExtendedGraph) BBHelper(node *model.Node, path []*model.Node, weight float64) {
	path = append(path, node)

	remaining := float64(len(g.Nodes)-len(path)+1) * 0.24
	if weight+remaining < *bestBBWeight {
		if len(path) < len(g.Nodes) {
			for _, edge := range node.Edges {
				if !g.contains(path, edge.Target) {
					g.BBHelper(g.Nodes[edge.Target], path, weight+edge.Weight)
				}
			}
		} else {
			lastEdge, _ := g.getEdge(path[len(path)-1].ID, path[0].ID)
			weight += lastEdge.Weight

			if weight < *bestBBWeight {
				bestWeight := weight
				bestPath := path

				fmt.Println("= New best path =")
				pathString := ""
				for _, node := range bestPath {
					pathString += fmt.Sprintf("%d -> ", node.ID)
				}
				fmt.Println(pathString[:len(pathString)-4])

				// bestBBPath = bestPath
				bestBBWeight = &bestWeight
			}
		}
	}
}
