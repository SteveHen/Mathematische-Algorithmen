package algorithms

import (
	"fmt"
	"math"

	"github.com/falo2/ma/model"
)

func (g *ExtendedGraph) CopyFullGraph() *ExtendedGraph {
	copy := &ExtendedGraph{}

	copy.Nodes = make(map[int64]*model.Node)
	for i := 0; i < len(g.Nodes); i++ {
		copy.Nodes[int64(i)] = &model.Node{ID: int64(i)}
	}

	for i := 0; i < len(g.Nodes); i++ {
		for j := 0; j < len(g.Nodes[int64(i)].Edges); j++ {
			edge := g.Nodes[int64(i)].Edges[j]
			copy.Nodes[edge.Source].Edges = append(copy.Nodes[edge.Source].Edges, &model.Edge{
				Source:   edge.Source,
				Target:   edge.Target,
				Weight:   edge.Weight,
				Capacity: edge.Capacity,
				Cost:     edge.Cost,
			})
		}
	}

	return copy
}

func (g *ExtendedGraph) MooreBellmanFordLegacy(source, target int64, verbose bool) (*ExtendedGraph, float64, []int64) {
	dist := make(map[int64]float64)
	previous := make(map[int64]*model.Node)
	pathWeight := 0.0
	result := &ExtendedGraph{Nodes: make(map[int64]*model.Node)}

	for i := 0; i < len(g.Nodes); i++ {
		dist[int64(i)] =
			float64(int64(^uint64(0) >> 1))
		result.Nodes[int64(i)] = &model.Node{ID: int64(i)}
	}

	dist[source] = 0
	previous[source] = g.Nodes[source]

	for i := 0; i < len(g.Nodes)-1; i++ {
		for _, node := range g.Nodes {
			for _, edge := range node.Edges {
				if dist[edge.Target] > dist[edge.Source]+edge.Weight {
					dist[edge.Target] = dist[edge.Source] + edge.Weight
					previous[edge.Target] = g.Nodes[edge.Source]
				}
			}
		}
	}

	// Check for negative cycles
	var cycle []int64
	for node := range g.Nodes {
		for _, edge := range g.Nodes[node].Edges {
			if dist[edge.Target] > dist[edge.Source]+edge.Weight {
				source := edge.Source
				visited := make(map[int64]bool)

				for !visited[source] {
					visited[source] = true
					cycle = append(cycle, source)
					source = previous[source].ID
				}

				cycle = append(cycle, source)
				return nil, -1, cycle
			}
		}
	}

	for node := 0; node < len(result.Nodes); node++ {
		nodeIndex := int64(node)
		if previous[nodeIndex] != nil {
			weight := dist[nodeIndex] - dist[previous[nodeIndex].ID]

			previousNode := previous[nodeIndex].ID
			result.Nodes[previousNode].Edges = append(result.Nodes[previousNode].Edges, &model.Edge{Source: previousNode, Target: nodeIndex, Weight: weight})
		}
	}

	if verbose {
		fmt.Println("\n= Distances =")
		for i := 0; i < len(dist); i++ {
			fmt.Println(i, dist[int64(i)])
		}
		fmt.Println()
	}

	pathWeight = math.Round(dist[target]*1e6) / 1e6
	fmt.Println(pathWeight)
	return result, pathWeight, nil
}

func (g *ExtendedGraph) EdmondsKarpLegacy(sourceNode, targetNode int64) (float64, *ExtendedGraph, error) {
	source := g.Nodes[sourceNode]
	target := g.Nodes[targetNode]

	if source == nil || target == nil {
		return 0, nil, fmt.Errorf("source or target node does not exist")
	}

  result := copyGraph(g)

	var maxFlow float64

	for {
		path, flow := result.BFSLegacy(source.ID, target.ID)

		if len(path) == 0 {
			break
		}

		maxFlow += flow

		for i := 0; i < len(path)-1; i++ {
			result.AddEdge(path[i], path[i+1], -flow)
			result.AddEdge(path[i+1], path[i], flow)
		}
	}

	maxFlow = math.Round(maxFlow*1e6) / 1e6
	return maxFlow, result, nil
}

func (g *ExtendedGraph) AddEdge(source int64, target int64, weight float64) {
	for _, edge := range g.Nodes[source].Edges {
		if edge.Target == target {
			edge.Weight += weight
			return
		}
	}

	g.Nodes[source].Edges = append(g.Nodes[source].Edges, &model.Edge{Source: source, Target: target, Weight: weight})
}

func (g *ExtendedGraph) BFSLegacy(source, target int64) ([]int64, float64) {
	visited := make(map[int64]bool)
	previous := make(map[int64]*model.Node)
	path := make([]int64, 0)
	flow := math.Inf(1)

	queue := make([]*model.Node, 0)
	queue = append(queue, g.Nodes[source])

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		visited[node.ID] = true

		if node.ID == target {
			path = append(path, node.ID)

			for previous[node.ID] != nil {
				path = append(path, previous[node.ID].ID)
				node = previous[node.ID]
			}

			// sort.Slice(path, func(i, j int) bool {
			// 	return path[i] < path[j]
			// })

			pathReverted := make([]int64, 0)
			for i := len(path)-1; i >= 0 ; i-- {
				pathReverted = append(pathReverted, path[i])
			}

			return pathReverted, flow
		}

		for _, edge := range node.Edges {
			if !visited[edge.Target] && edge.Weight > 0 {
				previous[edge.Target] = node
				queue = append(queue, g.Nodes[edge.Target])

				if edge.Weight < flow {
					flow = edge.Weight
				}
			}
		}
	}

	return path, flow
}
