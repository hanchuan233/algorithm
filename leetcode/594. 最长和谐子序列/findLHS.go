package main

import "fmt"

func findLHS(nums []int) int {
	record := make(map[int]int)
	max := 0
	for _, val := range nums {
		record[val]++
		if record[val-1] != 0 && record[val]+record[val-1] > max {
			max = record[val] + record[val-1]
		}
		if record[val+1] != 0 && record[val]+record[val+1] > max {
			max = record[val] + record[val+1]
		}
	}

	return max
}

func main() {
	nums := []int{1, 3, 2, 2, 5, 2, 3, 7}
	ret := findLHS(nums)
	fmt.Println(ret)
}
