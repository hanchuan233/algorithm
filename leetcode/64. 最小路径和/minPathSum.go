package main

import "fmt"

func minPathSum(grid [][]int) int {
	// dp[i] 表示在某一行时到达第i个点的代价，且覆盖上一行

	m := len(grid)
	n := len(grid[0])

	dp := make([]int, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i-1 >= 0 && j-1 >= 0 {
				dp[j] = Min(dp[j-1], dp[j])
			} else if i-1 >= 0 {
				dp[j] = dp[j]
			} else if j-1 >= 0 {
				dp[j] = dp[j-1]
			}

			dp[j] += grid[i][j]
		}
	}

	return dp[n-1]
}

func Min(i, j int) int {
	if i < j {
		return i
	} else {
		return j
	}
}

func main() {
	grid := [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}
	ret := minPathSum(grid)
	fmt.Println(ret)
}
