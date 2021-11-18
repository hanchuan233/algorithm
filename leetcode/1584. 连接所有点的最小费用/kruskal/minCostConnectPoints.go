package kruskal

import (
	"sort"
)

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

type unionFind struct {
	parent []int
	rank   []int
}

func NewUnionFind(n int) *unionFind {
	uf := unionFind{
		parent: make([]int, n),
		rank:   make([]int, n),
	}

	for idx := range uf.parent {
		uf.parent[idx] = idx
		uf.rank[idx] = idx
	}

	return &uf
}

func (uf *unionFind) Find(x int) int {
	if uf.parent[x] == x {
		return x
	} else {
		return uf.Find(uf.parent[x])
	}
}

func (uf *unionFind) Union(x, y int) bool {
	xFather, yFather := uf.Find(x), uf.Find(y)
	if xFather == yFather {
		return false
	}

	if uf.rank[xFather] < uf.rank[yFather] {
		uf.parent[xFather] = yFather
	} else if uf.rank[xFather] > uf.rank[yFather] {
		uf.parent[yFather] = xFather
	} else {
		uf.parent[xFather] = yFather
		uf.rank[yFather] += 1
	}
	return true
}

func MinCostConnectPoints(points [][]int) int {
	n := len(points)
	type edge struct {
		p1   int
		p2   int
		dist int
	}

	// 生成边
	var edges []edge
	for i, point := range points {
		for j := i + 1; j < n; j++ {
			edges = append(edges, edge{
				p1:   i,
				p2:   j,
				dist: dist(point, points[j]),
			})
		}
	}

	// 排序
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].dist < edges[j].dist
	})

	sum := 0
	uf := NewUnionFind(n)
	for _, edge := range edges {
		if uf.Union(edge.p1, edge.p2) {
			sum += edge.dist
			n--
			if n == 1 {
				break
			}
		}
	}

	return sum
}

