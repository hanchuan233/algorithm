package main

import "fmt"

// https://leetcode-cn.com/problems/single-number-iii/
func singleNumber(nums []int) []int {
	xorsum := 0
	for _, v := range nums {
		xorsum ^= v
	}

	group1 := 0
	group2 := 0
	k := xorsum & (-xorsum)
	for _, v := range nums {
		if v & k != 0 {
			group1 ^= v
		} else {
			group2 ^= v
		}
	}

	return []int{group1, group2}
}

func main() {
	nums := []int{1, 1, 2, 2, 3, 5}

	fmt.Println(singleNumber(nums))
}
