package main

import "fmt"

func colorBorder(grid [][]int, row int, col int, color int) [][]int {
	m := len(grid)
	n := len(grid[0])
	isBorder := func(r, c int) bool {
		if r == 0 || c == 0 || r == m-1 || c == n-1 {
			return true
		}

		if grid[r][c] != grid[r-1][c] || grid[r][c] != grid[r+1][c] || grid[r][c] != grid[r][c-1] || grid[r][c] != grid[r][c+1] {
			return true
		} else {
			return false
		}
	}

	type key struct {
		r, c int
	}

	seen := make(map[key]int)
	var ret [][]int
	var depthSearch func(r, c int)
	depthSearch = func(r, c int) {
		if r < 0 || c < 0 || r >= m || c >= n {
			return
		}
		if _, has := seen[key{r: r, c: c}]; has {
			return
		}
		if grid[r][c] != grid[row][col] {
			return
		}
		seen[key{r: r, c: c}] = 1
		if isBorder(r, c) {
			ret = append(ret, []int{r, c})
		}
		depthSearch(r+1, c)
		depthSearch(r-1, c)
		depthSearch(r, c+1)
		depthSearch(r, c-1)
	}

	depthSearch(row, col)
	for i := 0; i < len(ret); i++ {
		grid[ret[i][0]][ret[i][1]] = color
	}
	return grid
}

func main() {
	grid := [][]int{{1, 1}, {1, 2}}
	ret := colorBorder(grid, 0, 0, 3)
	fmt.Println(ret)
}
