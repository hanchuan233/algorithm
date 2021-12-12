package main

import (
	"fmt"
	"sort"
)

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)

	n := len(nums)
	if n == 1 {
		return [][]int{nums}
	}

	var ret [][]int
	var backtrack func(idx int)
	backtrack = func(idx int) {
		if idx == n {
			ret = append(ret, append([]int(nil), nums...))
			return
		}
		//tmp := append([]int(nil), nums[idx:]...)
		defer func() {
			//nums = append(nums[:idx], tmp...)
			nums = append(nums, nums[idx])
			nums = append(nums[:idx], nums[idx+1:]...)
		}()
		for i := idx; i < n; i++ {
			if i > idx && nums[i] == nums[idx] {
				continue
			}
			nums[i], nums[idx] = nums[idx], nums[i]
			backtrack(idx + 1)
		}
	}

	backtrack(0)
	return ret
}

func main() {
	nums := []int{2, 2, 1, 1}
	ret := permuteUnique(nums)
	fmt.Println(ret)
}
