package main

import "fmt"

func searchRange(nums []int, target int) []int {
	ret := []int{-1, -1}

	var getTargetIdx func(left, right int)
	getTargetIdx = func(left, right int) {
		for left <= right {
			mid := (left + right) / 2
			if nums[mid] == target {
				if ret[0] == -1 || mid < ret[0] {
					ret[0] = mid
				}
				if ret[1] == -1 || mid > ret[1] {
					ret[1] = mid
				}

				getTargetIdx(left, mid-1)
				getTargetIdx(mid+1, right)
				break
			} else if nums[mid] < target {
				left = mid + 1
			} else if nums[mid] > target {
				right = mid - 1
			}
		}
	}

	getTargetIdx(0, len(nums)-1)

	return ret
}

func main() {
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 8
	ret := searchRange(nums, target)
	fmt.Println(ret)
}
