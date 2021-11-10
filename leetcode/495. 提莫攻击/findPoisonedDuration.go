package main

import "fmt"

func findPoisonedDuration(timeSeries []int, duration int) int {
	startTime := -1
	endTime := -1
	totalTime := 0
	for _, val := range timeSeries {
		if startTime == -1 {
			startTime = val
			endTime = startTime + duration
		} else if val > endTime {
			// 结算一次
			totalTime += endTime - startTime
			startTime = val
			endTime = startTime + duration
		} else {
			endTime = val + duration
		}
	}

	totalTime += endTime - startTime
	return totalTime
}

func main() {
	timeSeries := []int{1, 4}
	duration := 2
	ret := findPoisonedDuration(timeSeries, duration)
	fmt.Println(ret)
}
