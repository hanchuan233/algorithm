package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Solution struct {
	nums []int
}


func Constructor(nums []int) Solution {
	return Solution{nums: nums}
}


func (this *Solution) Reset() []int {
	return this.nums
}


func (this *Solution) Shuffle() []int {
	tmp := append([]int{}, this.nums...)
	n := len(tmp)
	rand.Seed(time.Now().Unix())
	for i := n - 1; i > 0; i-- {
		random := rand.Intn(i)
		tmp[i], tmp[random] = tmp[random], tmp[i]
	}
	return tmp
}

func main() {
	nums := []int{1, 2, 3}
	obj := Constructor(nums)
	ret := obj.Shuffle()
	fmt.Println(ret)
	ret = obj.Reset()
	fmt.Println(ret)
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
