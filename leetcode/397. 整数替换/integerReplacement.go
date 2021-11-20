package main

import "fmt"

func integerReplacement(n int) int {
	// dp[n] 表示数字变为n所需的最小替换次数
	// dp[n] = min(dp[n*2] + 1, dp[n+1] + 1, dp[n-1] + 1)
	// dp[1] = min(dp[2] + 1, dp[2] + 1, dp[0] + 1)

	var getCount func(i int, cnt int) int
	getCount = func(i int, cnt int) int {
		if i == 1 {
			return cnt
		}
		cnt++
		if i % 2 == 0 {
			return getCount(i / 2, cnt)
		} else {
			return min(getCount(i + 1, cnt), getCount(i - 1, cnt))
		}
	}
	return getCount(n, 0)
}

func min(i, j int) int {
	if i < j {
		return i
	} else {
		return j
	}
}

func main()  {
	ret := integerReplacement(8)
	fmt.Println(ret)
}