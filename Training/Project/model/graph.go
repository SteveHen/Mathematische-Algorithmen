package model

type Graph struct {
	Nodes map[int64]*Node
}

type Node struct {
	ID    int64
	Edges []*Edge
}

type Edge struct {
	Source   int64
	Target   int64
	Weight   float64
	Capacity float64
	Cost     float64
	Residual bool
}
