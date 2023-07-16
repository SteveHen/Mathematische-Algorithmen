package datatypes

import (
	"container/heap"

	"github.com/falo2/ma/model"
)

type PriorityQueue []model.Edge

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Weight < pq[j].Weight
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(edge interface{}) {
	*pq = append(*pq, edge.(model.Edge))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	edge := old[n-1]
	*pq = old[0 : n-1]
	return edge
}

func (pq *PriorityQueue) Insert(edge model.Edge) {
	heap.Push(pq, edge)
}

func (pq *PriorityQueue) IsEmpty() bool {
	return len(*pq) == 0
}

func (pq *PriorityQueue) Peek() model.Edge {
	return (*pq)[0]
}
