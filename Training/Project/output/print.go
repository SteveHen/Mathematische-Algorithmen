package output

import (
	"fmt"
	"strconv"

	"github.com/falo2/ma/algorithms"
)

func Print(g *algorithms.ExtendedGraph, isWeighted, isBalanced bool) {
	if isBalanced {
		for i := 0; i < len(g.Nodes); i++ {
			fmt.Printf("Node %s -> ", strconv.FormatInt(g.Nodes[int64(i)].ID, 10))
			for _, value2 := range g.Nodes[int64(i)].Edges {
				fmt.Printf("%s ", strconv.FormatInt(value2.Target, 10))
				fmt.Printf("(Fl: %s, Ca: %s, Co: %s) | ", strconv.FormatFloat(value2.Weight, 'f', 6, 64), strconv.FormatFloat(value2.Capacity, 'f', 6, 64), strconv.FormatFloat(value2.Cost, 'f', 6, 64))
			}
			fmt.Println()
		}
		return
	}

	if isWeighted {
		for i := 0; i < len(g.Nodes); i++ {
			fmt.Printf("Node %s -> ", strconv.FormatInt(g.Nodes[int64(i)].ID, 10))
			for _, value2 := range g.Nodes[int64(i)].Edges {
				fmt.Printf("%s ", strconv.FormatInt(value2.Target, 10))
				fmt.Printf("(W: %s) | ", strconv.FormatFloat(value2.Weight, 'f', 6, 64))
			}
			fmt.Println()
		}
		return
	}

	for i := 0; i < len(g.Nodes); i++ {
		fmt.Printf("Node %s -> ", strconv.FormatInt(g.Nodes[int64(i)].ID, 10))
		for _, value2 := range g.Nodes[int64(i)].Edges {
			fmt.Printf("%s | ", strconv.FormatInt(value2.Target, 10))
		}
		fmt.Println()
	}
}
