package main

import "fmt"

func grayCode(n int) []int {
	// dp[n] = dp[n-1], reverse(dp[n-1] + 2^(n-1))
	// dp[n] = dp[n-1], reverse(dp[n-1] + 1<<(n-1))
	ret := make([]int, 1<<n)
	for i := 1; i <= n; i++ {
		for j := 1 << (i - 1); j < 1<<i; j++ {
			ret[j] = ret[1<<i-j-1] + 1<<(i-1)
		}
	}

	return ret
}

func main() {
	fmt.Println(grayCode(2))
}
