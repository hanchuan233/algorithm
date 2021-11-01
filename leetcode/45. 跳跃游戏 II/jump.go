package main

import "fmt"

func jump(nums []int) int {
	// curStep := len(nums) - 1
	// count := 0
	// for {
	//     if curStep == 0 {
	//         return count
	//     }
	//     for i := 0; i < curStep; i++ {
	//         if i + nums[i] >= curStep {
	//             curStep = i
	//             count++
	//         }
	//     }
	// }
	farthest := 0
	curFarthest := 0
	count := 0
	for i := 0; i < len(nums)-1; i++ {
		farthest = max(farthest, i+nums[i])
		if i == curFarthest {
			curFarthest = farthest
			count++
		}
	}

	return count
}

func max(i, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}

func main() {
	nums := []int{2, 3, 1, 1, 4}
	ret := jump(nums)
	fmt.Println(ret)
}
