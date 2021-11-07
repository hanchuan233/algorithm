package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)

	ret := make([][]int, 0)
	for i := 0; i < len(nums)-2; i++ {
		if i >= 1 && nums[i-1] == nums[i] {
			continue
		}
		j := i + 1
		k := len(nums) - 1
		for ; j < k; {
			if nums[i]+nums[j]+nums[k] < 0 {
				j++
			} else if nums[i]+nums[j]+nums[k] > 0 {
				k--
			} else if nums[i]+nums[j]+nums[k] == 0 {
				ret = append(ret, []int{nums[i], nums[j], nums[k]})
				for ; j < len(nums)-1; j++ {
					if nums[j] != nums[j+1] {
						j++
						break
					}
				}

				for ; k > i+1; k-- {
					if nums[k] != nums[k-1] {
						k--
						break
					}
				}
			}
		}
	}

	return ret
}

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	ret := threeSum(nums)
	fmt.Println(ret)
}
