package input

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/falo2/ma/algorithms"
	"github.com/falo2/ma/model"
)

func Read(path string, isDirected, hasBalance bool) (*algorithms.ExtendedGraph, *map[int64]float64, bool) {
	newGraph := algorithms.ExtendedGraph{}
	newGraph.Nodes = make(map[int64]*model.Node)
	file, _ := os.ReadFile(path)

	fileContent := bytes.Split(file, []byte("\n"))
	nodes, err := strconv.Atoi(strings.TrimSpace(string(fileContent[0])))
	if err != nil {
		return nil, nil, false
	}

	for i := 0; i < nodes; i++ {
		newGraph.Nodes[int64(i)] = &model.Node{ID: int64(i)}
	}

	var isWeighted bool
	edgePosition := 1
	balanceString := ""
	balance := make(map[int64]float64)
	if hasBalance {
		for key, value := range fileContent[1 : nodes+1] {
			balanceString = strings.TrimSpace(string(value))
			balance[int64(key)], _ = strconv.ParseFloat(balanceString, 64)
		}

		edgePosition = nodes + 1
	}

	for _, content := range fileContent[edgePosition:] {
		properties := strings.Join(strings.Fields(string(content)), " ")
		propertySlice := strings.Split(properties, " ")

		if len(propertySlice) > 1 {
			source := propertySlice[0]
			sourceID, _ := strconv.ParseInt(source, 10, 64) // convert the string to int64
			target := propertySlice[1]
			targetID, _ := strconv.ParseInt(target, 10, 64)

			var newSourceEdge, newTargetEdge model.Edge

			if len(propertySlice) == 2 {
				newSourceEdge = model.Edge{Source: sourceID, Target: targetID, Weight: 0}
				if !isDirected {
					newTargetEdge = model.Edge{Source: targetID, Target: sourceID, Weight: 0}
				}

				newGraph.Nodes[sourceID].Edges = append(newGraph.Nodes[sourceID].Edges, &newSourceEdge)

				if !isDirected {
					newGraph.Nodes[targetID].Edges = append(newGraph.Nodes[targetID].Edges, &newTargetEdge)
				}
				continue
			}

			if len(propertySlice) > 2 {
				weightNum := 0.0
				capacityNum := 0.0

				if len(propertySlice) >= 3 {
					weight := propertySlice[2]
					weightNum, _ = strconv.ParseFloat(weight, 64)

					isWeighted = true
				}

				if len(propertySlice) == 3 {
					newSourceEdge = model.Edge{Source: sourceID, Target: targetID, Weight: weightNum}
					if !isDirected {
						newTargetEdge = model.Edge{Source: targetID, Target: sourceID, Weight: weightNum}
					}
					newGraph.Nodes[sourceID].Edges = append(newGraph.Nodes[sourceID].Edges, &newSourceEdge)

					if !isDirected {
						newGraph.Nodes[targetID].Edges = append(newGraph.Nodes[targetID].Edges, &newTargetEdge)
					}
					continue
				}

				if len(propertySlice) == 4 {
					capacity := propertySlice[3]
					capacityNum, _ = strconv.ParseFloat(capacity, 64)

					newSourceEdge = model.Edge{Source: sourceID, Target: targetID, Weight: 0, Capacity: capacityNum, Cost: weightNum}
					newGraph.Nodes[sourceID].Edges = append(newGraph.Nodes[sourceID].Edges, &newSourceEdge)

					if !isDirected {
						newTargetEdge = model.Edge{Source: targetID, Target: sourceID, Weight: 0, Capacity: capacityNum, Cost: weightNum}
						newGraph.Nodes[targetID].Edges = append(newGraph.Nodes[targetID].Edges, &newTargetEdge)
					}
					continue
				}
			}
		}
	}

	fmt.Println(strconv.Itoa(len(newGraph.Nodes)) + " nodes read from file.")
	return &newGraph, &balance, isWeighted
}
