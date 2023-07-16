package algorithms

import (
	"fmt"
	"math"
	"sort"

	"github.com/falo2/ma/model"
)

func (g *ExtendedGraph) Kruskal() (*ExtendedGraph, []*model.Edge, string) {
	totalWeight := 0.0
	if len(g.Nodes) == 0 {
		return nil, nil, ""
	}

	var edges []*model.Edge
	for _, n := range g.Nodes {
		edges = append(edges, n.Edges...)
	}
	sort.Slice(edges, func(i, j int) bool { return edges[i].Weight < edges[j].Weight })

  var resultEdges []*model.Edge

	mst := &ExtendedGraph{Nodes: make(map[int64]*model.Node)}
	for _, e := range edges {
		if _, ok := mst.Nodes[e.Source]; !ok {
			mst.Nodes[e.Source] = &model.Node{ID: e.Source}
		}
		if _, ok := mst.Nodes[e.Target]; !ok {
			mst.Nodes[e.Target] = &model.Node{ID: e.Target}
		}
		if g.find(mst.Nodes[e.Source]) != g.find(mst.Nodes[e.Target]) {
			g.union(mst.Nodes[e.Source], mst.Nodes[e.Target])
			totalWeight += e.Weight
			mst.Nodes[e.Source].Edges = append(mst.Nodes[e.Source].Edges, e)
			mst.Nodes[e.Target].Edges = append(mst.Nodes[e.Target].Edges, &model.Edge{Source: e.Target, Target: e.Source, Weight: e.Weight})

      resultEdges = append(resultEdges, e)
		}
	}

	totalWeight = math.Round(totalWeight*1e6) / 1e6
	return mst, resultEdges, fmt.Sprintf("The MST contains edges with a total weight of %s.", fmt.Sprint(totalWeight))
}

func (g *ExtendedGraph) find(node *model.Node) *model.Node {
	if node.ID != g.Nodes[node.ID].ID {
		g.Nodes[node.ID] = g.find(g.Nodes[node.ID])
	}
	return g.Nodes[node.ID]
}

func (g *ExtendedGraph) union(node1, node2 *model.Node) {
	root1 := g.find(node1)
	root2 := g.find(node2)
	g.Nodes[root2.ID] = root1
}
