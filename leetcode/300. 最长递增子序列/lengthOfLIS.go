package main

import "fmt"

func lengthOfLIS(nums []int) int {
	// tail[i] 存储表示长度为i的最长子序列的最后一个值，且该数组为单调递增
	var tail []int
	for _, v := range nums {
		if len(tail) == 0 || v > tail[len(tail)-1] {
			tail = append(tail, v)
			continue
		}
		left := 0
		right := len(tail) - 1
		for left <= right {
			mid := (left + right) / 2
			if v == tail[mid] {
				break
			} else if v < tail[mid] {
				if mid > 0 {
					if v > tail[mid-1] {
						tail[mid] = v
						break
					} else {
						right = mid - 1
					}
				} else {
					tail[mid] = v
					break
				}
			} else {
				left = mid + 1
			}
		}
	}

	return len(tail)
}

func main() {
	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	ret := lengthOfLIS(nums)
	fmt.Println(ret)
}
