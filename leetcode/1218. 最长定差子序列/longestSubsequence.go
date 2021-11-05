package main

import "fmt"

func longestSubsequence(arr []int, difference int) int {
    max := 1

    mp := make(map[int]int, len(arr))

    for _, val := range arr {
        if _, has := mp[val - difference]; has {
            mp[val] = mp[val - difference] + 1
        } else {
            mp[val] = 1
        }

        if mp[val] > max {
            max = mp[val]
        }
    }

    return max
}

func main(){
    arr := []int{1, 2, 3, 4}
    difference := 1
    ret := longestSubsequence(arr, difference)
    fmt.Println(ret)
}

