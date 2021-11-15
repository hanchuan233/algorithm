package main

import "fmt"

func maxProfit(prices []int) int {
	// 贪心
	// profit := 0
	// for i := 1; i < len(prices); i++ {
	//     if prices[i] > prices[i - 1] {
	//         profit += prices[i] - prices[i - 1]
	//     }
	// }
	// return profit

	// 动态规划
	// dp[0][n] 表示在第n天不持有股票的情况下的收益
	// dp[1][n] 表示在第n天持有股票的情况下的收益
	// dp[0][n] = max(dp[0][n-1], dp[1][n-1] + prices[n])
	// dp[1][n] = max(dp[1][n-1], dp[0][n-1] - prices[n])

	dp := make([][]int, 2)
	dp[0] = make([]int, len(prices))
	dp[1] = make([]int, len(prices))

	dp[0][0] = 0
	dp[1][0] = -prices[0]
	for i := 1; i < len(prices); i++ {
		dp[0][i] = max(dp[0][i-1], dp[1][i-1]+prices[i])
		dp[1][i] = max(dp[1][i-1], dp[0][i-1]-prices[i])
	}

	return dp[0][len(prices)-1]
}

func max(i, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	ret := maxProfit(prices)
	fmt.Println(ret)
}
