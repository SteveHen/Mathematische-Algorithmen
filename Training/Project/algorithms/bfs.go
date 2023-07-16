package algorithms

import (
	"github.com/falo2/ma/model"
)

func (g *ExtendedGraph) BFS(source, target int64, visited map[int64]bool) ([]*model.Node, map[int64]bool) {
	result := []*model.Node{}
	queue := []int64{source}
	previous := make(map[int64]int64)

	if visited == nil {
		visited = make(map[int64]bool)
	}

	visited[source] = true
	previous[source] = source

	for len(queue) > 0 {
		currentNode := queue[0]
		queue = queue[1:]

		for i := 0; i < len(g.Nodes[currentNode].Edges); i++ {
			edge := g.Nodes[currentNode].Edges[i]
			currentTarget := edge.Target

			if !visited[currentTarget] && edge.Capacity > 0 {
				visited[currentTarget] = true
				queue = append(queue, currentTarget)
				previous[currentTarget] = currentNode

				if currentTarget == target {
					tmp := currentTarget

					for tmp != source {
						result = append([]*model.Node{g.Nodes[tmp]}, result...)
						tmp = previous[tmp]
					}

					result = append([]*model.Node{g.Nodes[source]}, result...)
					return result, visited
				}
			}
		}
	}
	return result, visited
}
