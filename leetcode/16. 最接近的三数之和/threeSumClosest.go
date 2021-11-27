package main

import (
	"fmt"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	n := len(nums)
	sum := nums[0] + nums[1] + nums[2]
	delta := Abs(target - sum)
	for i := 0; i < n-2; i++ {
		j := i + 1
		k := n - 1
		for j < k {
			curSum := nums[i] + nums[j] + nums[k]

			if Abs(target-curSum) < delta {
				sum = curSum
				delta = Abs(target - sum)
				if delta == 0 {
					return sum
				}
			}

			if curSum < target {
				j++
			} else {
				k--
			}
		}
	}

	return sum
}

func Abs(i int) int {
	if i < 0 {
		return -i
	} else {
		return i
	}
}

func main() {
	nums := []int{-1, 2, 1, -4}
	target := 1
	ret := threeSumClosest(nums, target)
	fmt.Println(ret)
}
