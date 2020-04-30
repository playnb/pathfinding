package pathfinding

import (
	"container/heap"
	//"fmt"
)

type PriorityQueue []*Node

// heap.Interface
func (pq PriorityQueue) Len() int {
	return len(pq)
}

// heap.Interface
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority() < pq[j].priority()
}

// heap.Interface
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].heapIndex = i
	pq[j].heapIndex = j
}

// heap.Interface
func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Node)
	item.heapIndex = n
	*pq = append(*pq, item)
}

// heap.Interface
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.heapIndex = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

/* Node */

func (pq *PriorityQueue) PushNode(n *Node) {
	heap.Push(pq, n)
}

func (pq *PriorityQueue) PopNode() *Node {
	return heap.Pop(pq).(*Node)
}

func (pq *PriorityQueue) RemoveNode(n *Node) {
	heap.Remove(pq, n.heapIndex)
}

func (pq *PriorityQueue) UpdateNode(n *Node) {
	heap.Fix(pq, n.heapIndex)
}
