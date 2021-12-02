package main

import "fmt"

// func trap(height []int) int {
//     n := len(height)
//     sum := 0
//     for i := 0; i < n; i++ {
//         leftMax := 0
//         rightMax := 0
//         for j := i - 1; j >= 0; j-- {
//             leftMax = max(leftMax, height[j])
//         }
//         for k := i + 1; k < n; k++ {
//             rightMax = max(rightMax, height[k])
//         }
//         curHeight := min(leftMax, rightMax) - height[i]
//         if curHeight > 0 {
//             sum += curHeight
//         }
//     }

//     return sum
// }

func trap(height []int) int {
	n := len(height)

	leftMax := make([]int, n)
	rightMax := make([]int, n)

	curMax := 0
	// 从左遍历 得到i对应的leftMax
	for i := 0; i < n; i++ {
		leftMax[i] = curMax
		curMax = max(curMax, height[i])
	}

	curMax = 0
	// 从右遍历 得到i对应的rightMax
	for i := n - 1; i >= 0; i-- {
		rightMax[i] = curMax
		curMax = max(curMax, height[i])
	}

	sum := 0
	for i, v := range height {
		curHeight := min(leftMax[i], rightMax[i])
		if curHeight > v {
			sum += curHeight - v
		}
	}
	return sum
}

func max(i, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}

func min(i, j int) int {
	if i < j {
		return i
	} else {
		return j
	}
}

func main() {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	ret := trap(height)
	fmt.Println(ret)
}
