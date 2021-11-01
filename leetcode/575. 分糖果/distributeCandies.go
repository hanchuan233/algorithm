package main

import "fmt"

func distributeCandies(candyType []int) int {
	set := make(map[int]int)
	for _, v := range candyType {
		set[v] = 1
	}

	return Min(len(candyType)/2, len(set))
}

func Min(i, j int) int {
	if i < j {
		return i
	} else {
		return j
	}
}

func main() {
	candyType := []int{1, 1, 2, 2, 3, 3}
	ret := distributeCandies(candyType)
	fmt.Println(ret)
}
