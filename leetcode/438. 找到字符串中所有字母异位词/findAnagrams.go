package main

import "fmt"

func findAnagrams(s string, p string) []int {
	sLen := len(s)
	pLen := len(p)
	if sLen < pLen {
		return []int{}
	}

	sCount := [26]int{}
	pCount := [26]int{}

	for i, v := range p {
		sCount[s[i] - 'a']++
		pCount[v - 'a']++
	}

	var ret []int
	if sCount == pCount {
		ret = append(ret, 0)
	}

	for i := 1; i <= len(s) - pLen; i++ {
		sCount[s[i-1]-'a']--
		sCount[s[i+pLen-1]-'a']++

		if sCount == pCount {
			ret = append(ret, i)
		}
	}

	return ret
}

func main() {
	s := "cbaebabacd"
	p := "abc"
	ret := findAnagrams(s, p)
	fmt.Println(ret)
}
