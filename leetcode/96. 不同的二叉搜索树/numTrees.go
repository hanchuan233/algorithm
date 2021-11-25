package main

import "fmt"

func numTrees(n int) int {
	// dp[n] = dp[0] * dp[n-1] + dp[1] * dp[n-2] + ... + dp[n-1] * dp[0]
	dp := make([]int, n + 1)
	dp[0] = 1
	dp[1] = 1

	for i := 2; i <= n; i++ {
		for j := 0; j < i; j++ {
			dp[i] += dp[j] * dp[i-1-j]
		}
	}
	return dp[n]
}

func main()  {
	ret := numTrees(3)
	fmt.Println(ret)
}