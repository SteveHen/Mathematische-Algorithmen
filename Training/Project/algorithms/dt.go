package algorithms

import (
	"fmt"

	"github.com/falo2/ma/model"
)

func (g *ExtendedGraph) DoubleTreeAlgorithm() float64 {
	newGraph := copyGraph(g)
	mst, _, _ := newGraph.Kruskal()
	pathToTarget := mst.DFS(0, make(map[int64]bool), make([]*model.Node, 0))

	var (
		weight float64
		edge   *model.Edge
	)

	fmt.Println("\n= Path =")
	pathString := ""
	for _, node := range pathToTarget {
		pathString += fmt.Sprintf("%d -> ", node.ID)
	}
	fmt.Println(pathString[:len(pathString)-4])

	for i := 0; i < len(pathToTarget)-1; i++ {
		edge, _ = g.getEdge(pathToTarget[i].ID, pathToTarget[i+1].ID)
		weight += edge.Weight
	}

	lastEdge, _ := g.getEdge(pathToTarget[len(pathToTarget)-1].ID, pathToTarget[0].ID)
	weight += lastEdge.Weight

	fmt.Println("\n= Weight =")
	fmt.Println(weight)
	return weight
}

func (g *ExtendedGraph) getEdge(source, target int64) (*model.Edge, int) {
	if g.Nodes[source] != nil && len(g.Nodes[source].Edges) > 0 {
		for i := 0; i < len(g.Nodes[source].Edges); i++ {
			edge := g.Nodes[source].Edges[i]
			if edge.Source == source && edge.Target == target {
				return edge, i
			}
		}
	}

	return nil, 0
}

func copyGraph(g *ExtendedGraph) *ExtendedGraph {
	newGraph := new(ExtendedGraph)
	newGraph.Nodes = make(map[int64]*model.Node)

	for key, node := range g.Nodes {
		newGraph.Nodes[key] = node
	}

	return newGraph
}
