package main

import "fmt"

func maxSumOfThreeSubarrays(nums []int, k int) []int {
	n := len(nums)
	sum := make([]int, n-k+1)
	for i := 0; i < k; i++ {
		sum[0] += nums[i]
	}

	for i := 1; i < n-k; i++ {
		sum[i] = sum[i-1] - nums[i-1] + nums[i+k-1]
	}

	// dp[i][j] 表示前i个数所构成的j个无重叠子数组的最大和
	// dp[i][j] = max(dp[i-1][j], dp[i-k][j-1]+sum[i-k+1])
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 4)
	}
	for i := 0; i < n; i++ {
		for j := 1; j <= 3; j++ {
			if i < k*j-1 {
				break
			}
			if i > k {
				dp[i][j] = max(dp[i-1][j], dp[i-k][j-1]+sum[i-k+1])
			} else {
				dp[i][j] = max(dp[i-1][j], sum[i-k+1])
			}

		}
	}

	fmt.Println(dp[n-1][3])
	return []int{}
}

func max(i, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}

func main() {
	nums := []int{1, 2, 1, 2, 6, 7, 5, 1}
	k := 2
	ret := maxSumOfThreeSubarrays(nums, k)
	fmt.Println(ret)
}
