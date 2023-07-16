package algorithms

import (
	"fmt"

	"github.com/falo2/ma/model"

	"math"
)

func (g *ExtendedGraph) NearestNeighbor() []*model.Node {
	visited := make(map[int64]bool)
	var path []*model.Node
	totalWeight := 0.0

	start := g.Nodes[0]
	path = append(path, start)
	visited[start.ID] = true

	for len(visited) < len(g.Nodes) {
		current := path[len(path)-1]
		var next *model.Node
		var minDist = math.MaxFloat64

		for _, edge := range current.Edges {
			if !visited[edge.Target] && edge.Weight < minDist {
				next = g.Nodes[edge.Target]
				minDist = edge.Weight
			}
		}

		if next != nil {
			path = append(path, next)
			visited[next.ID] = true
			totalWeight += minDist
		}
	}

	for i := 0; i < len(path)-1; i++ {
		path[i].Edges = append(path[i].Edges, &model.Edge{
			Source: path[i].ID,
			Target: path[i+1].ID,
			Weight: 0,
		})
		/*path[i+1].Edges = append(path[i+1].Edges, &model.Edge{
			Source: path[i+1].ID,
			Target: path[i].ID,
			Weight: 0,
		}) */
	}

	for _, edge := range path[len(path)-1].Edges {
		if edge.Target == start.ID {
			totalWeight += edge.Weight
			break
		}
	}

	fmt.Println("\n= Path =")
	pathString := ""
	for _, node := range path {
		pathString += fmt.Sprintf("%d -> ", node.ID)
	}
	fmt.Println(pathString[:len(pathString)-4])

	fmt.Println("\n= Weight =")
	fmt.Println(totalWeight)
	return path
}
