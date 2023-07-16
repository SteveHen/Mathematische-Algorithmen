package algorithms

import (
	"github.com/falo2/ma/model"
)

type GraphInterface interface {
	DepthSearch(node int64)
}

type ExtendedGraph model.Graph

func (g *ExtendedGraph) DepthSearch(start, target int64) []*model.Node {
	visited := make(map[int64]bool)
	path := g.DFS(start, visited, nil)
	// countGraphs := 1

	// could delete this
	for key := range g.Nodes {
		if !visited[key] {
			path = g.DFS(key, visited, path)
			// countGraphs++
		}
	}

	// fmt.Printf("There exist %s correlation components.\n", strconv.Itoa(countGraphs))
	return path
}

func (g *ExtendedGraph) DFS(node int64, visited map[int64]bool, path []*model.Node) []*model.Node {
	// IMPORTANT
	if visited[node] {
		return path
	}

	visited[node] = true
	path = append(path, g.Nodes[node])

	for _, edge := range g.Nodes[node].Edges {
		var nextNode *model.Node
		if edge.Source == node {
			nextNode = g.Nodes[edge.Target]
		} else {
			nextNode = g.Nodes[edge.Source]
		}
		path = g.DFS(nextNode.ID, visited, path)
	}

	return path
}
