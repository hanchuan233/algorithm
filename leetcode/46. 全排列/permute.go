package main

import "fmt"

func permute(nums []int) [][]int {
	n := len(nums)
	if n == 1 {
		return [][]int{nums}
	}

	var ret [][]int
	var backtrack func(idx int)
	backtrack = func(idx int) {
		if idx >= n-1 {
			ret = append(ret, append([]int(nil), nums...))
			return
		}

		for i := idx; i < n; i++ {
			nums[i], nums[idx] = nums[idx], nums[i]
			backtrack(idx + 1)
			nums[i], nums[idx] = nums[idx], nums[i]
		}
	}

	backtrack(0)
	return ret
}

func main() {
	nums := []int{1, 2, 3}
	ret := permute(nums)
	fmt.Println(ret)
}
