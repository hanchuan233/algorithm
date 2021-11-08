package main

import "fmt"

type UnionFind struct {
	parents []int
	ranks   []int
	lands   int
}

func MakeUnionFind(grid [][]byte) *UnionFind {
	m := len(grid)
	n := len(grid[0])
	uf := &UnionFind{
		parents: make([]int, m*n),
		ranks:   make([]int, m*n),
		lands:   0,
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				uf.lands++
				uf.parents[i*n+j] = i*n + j
			}
		}
	}
	return uf
}

func (uf *UnionFind) Find(x int) int {
	if uf.parents[x] == x {
		return x
	} else {
		return uf.Find(uf.parents[x])
	}
}

func (uf *UnionFind) Union(x, y int) {
	xFather := uf.Find(x)
	yFather := uf.Find(y)
	if xFather == yFather {
		return
	}
	if uf.ranks[xFather] < uf.ranks[yFather] {
		uf.parents[xFather] = yFather
	} else if uf.ranks[xFather] > uf.ranks[yFather] {
		uf.parents[yFather] = xFather
	} else {
		uf.parents[yFather] = xFather
		uf.ranks[xFather] += 1
	}
	uf.lands--
}

func numIslands(grid [][]byte) int {
	m := len(grid)
	n := len(grid[0])

	uf := MakeUnionFind(grid)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				grid[i][j] = '0'
				if i-1 >= 0 && grid[i-1][j] == '1' {
					uf.Union(i*n+j, (i-1)*n+j)
				}
				if i+1 < m && grid[i+1][j] == '1' {
					uf.Union(i*n+j, (i+1)*n+j)
				}
				if j-1 >= 0 && grid[i][j-1] == '1' {
					uf.Union(i*n+j, i*n+(j-1))
				}
				if j+1 < n && grid[i][j+1] == '1' {
					uf.Union(i*n+j, i*n+(j+1))
				}
			}
		}
	}

	return uf.lands
}

func main() {
	grid := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}
	ret := numIslands(grid)
	fmt.Println(ret)
}
