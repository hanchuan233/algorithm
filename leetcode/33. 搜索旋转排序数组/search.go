package main

import "fmt"

func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		} else if nums[left] <= nums[mid] {
			// 左边有序
			if target >= nums[left] && target <= nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			// 右边有序
			if target >= nums[mid] && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return -1
}

func main() {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	target := 0
	ret := search(nums, target)
	fmt.Println(ret)
}
