package main

import "fmt"

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	// f(i, j) = f(i-1, j) + f(i, j-1)
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	dp := make([]int, n)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 {
				// 第一排
				if obstacleGrid[i][j] == 1 {
					dp[j] = 0
					break
				} else {
					if j == 0 {
						dp[j] = 1
					} else {
						dp[j] = dp[j-1]
					}
				}
			} else {
				// 不是第一排
				if obstacleGrid[i][j] == 1 {
					dp[j] = 0
				} else {
					if j != 0 {
						dp[j] = dp[j-1] + dp[j]
					}
				}
			}
		}
	}

	return dp[n-1]
}

func main() {
	obstacleGrid := [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}
	ret := uniquePathsWithObstacles(obstacleGrid)
	fmt.Println(ret)
}
