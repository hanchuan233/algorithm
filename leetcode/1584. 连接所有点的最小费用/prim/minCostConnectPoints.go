package prim

import "container/heap"

func abs(i int) int {
	if i >= 0 {
		return i
	} else {
		return -i
	}
}

func dist(p, q []int) int {
	return abs(p[0]-q[0]) + abs(p[1]-q[1])
}

type PriorityQueue []*Edge

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].dist < pq[j].dist }
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*Edge))
}

func (pq *PriorityQueue) Pop() interface{} {
	n := len(*pq)
	x := (*pq)[n-1]
	*pq = (*pq)[0 : n-1]
	return x
}

type Edge struct {
	p1   int
	p2   int
	dist int
}

func MinCostConnectPoints(points [][]int) int {
	n := len(points)
	isVisited := make([]bool, n)
	pq := PriorityQueue{}

	// 将当前点的所有的边放入优先级队列
	getEdges := func(idx int, point []int) {
		for i, p := range points {
			if !isVisited[i] && idx != i {
				heap.Push(&pq, &Edge{
					p1:   idx,
					p2:   i,
					dist: dist(p, point),
				})
			}
		}
	}

	// 返回值为未被访问过的点的序号 如果两个点都被访问过则返回-1
	checkVisited := func(edge Edge) int {
		if !isVisited[edge.p2] {
			return edge.p2
		}
		return -1
	}

	getEdges(0, points[0])
	isVisited[0] = true
	n--

	sum := 0
	for n > 0 {
		edge := heap.Pop(&pq).(*Edge)
		idx := checkVisited(*edge)
		if idx == -1 {
			continue
		}

		sum += edge.dist
		getEdges(idx, points[idx])
		isVisited[idx] = true
		n--
	}

	return sum
}