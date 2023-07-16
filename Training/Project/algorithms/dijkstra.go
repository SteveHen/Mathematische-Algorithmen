package algorithms

import (
	"fmt"
	"math"
	"strconv"

	"github.com/falo2/ma/model"
)

func (g *ExtendedGraph) Dijkstra(source, target int64, verbose bool) (*ExtendedGraph, float64) {
	dist := make(map[int64]float64)
	previous := make(map[int64]*model.Node)
	visited := make(map[int64]bool)
	pathWeight := 0.0

	result := &ExtendedGraph{Nodes: make(map[int64]*model.Node)}

	for i := 0; i < len(g.Nodes); i++ {
		dist[int64(i)] = math.Inf(1)
		result.Nodes[int64(i)] = &model.Node{ID: int64(i)}
	}

	dist[source] = 0
	previous[source] = g.Nodes[source]

	resultString := ""

	for i := 0; i < len(g.Nodes); i++ {
		minNode := minDistance(dist, visited)
		visited[minNode] = true

		resultString += fmt.Sprintf("%s -> ", strconv.Itoa(int(minNode)))

		if minNode == target {
			break
		}

		for _, edge := range g.Nodes[minNode].Edges {
			if !visited[edge.Target] && dist[edge.Target] > dist[minNode]+edge.Weight {
				dist[edge.Target] = dist[minNode] + edge.Weight
				previous[edge.Target] = g.Nodes[minNode]
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
		resultString = "\n= Path =\n" + resultString[:len(resultString)-4] + "\n"
		fmt.Println(resultString)

		fmt.Println("\n= Distances =")
		for i := 0; i < len(dist); i++ {
			fmt.Println(i, dist[int64(i)])
		}
		fmt.Println()
	}

	pathWeight = math.Round(dist[target]*1e6) / 1e6
	return result, pathWeight
}

func minDistance(dist map[int64]float64, visited map[int64]bool) int64 {
	min := math.Inf(1)
	minIndex := int64(-1)

	for i := 0; i < len(dist); i++ {
		if !visited[int64(i)] && dist[int64(i)] <= min {
			min = dist[int64(i)]
			minIndex = int64(i)
		}
	}

	return minIndex
}
